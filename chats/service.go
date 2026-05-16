package chats

import (
	"errors"
	"student-app/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePrivateChat(userID, otherUserID int) (*ChatResponse, error) {
	// Проверяем, существует ли уже приватный чат между пользователями
	// Для упрощения создаем новый

	chat := &models.Chat{
		Name: "",
		Type: "private",
	}

	if err := s.repo.CreateChat(chat); err != nil {
		return nil, err
	}

	members := []int{userID, otherUserID}
	if err := s.repo.AddMembers(chat.ID, members); err != nil {
		return nil, err
	}

	return s.getChatResponse(chat.ID, userID)
}

func (s *Service) CreateGroupChat(req *CreateChatRequest, creatorID int) (*ChatResponse, error) {
	chat := &models.Chat{
		Name: req.Name,
		Type: req.Type,
	}

	if chat.Type == "" {
		chat.Type = "group"
	}

	if err := s.repo.CreateChat(chat); err != nil {
		return nil, err
	}

	// Добавляем всех участников
	members := append(req.Members, creatorID)
	if err := s.repo.AddMembers(chat.ID, members); err != nil {
		return nil, err
	}

	return s.getChatResponse(chat.ID, creatorID)
}

func (s *Service) GetUserChats(userID int) ([]ChatResponse, error) {
	chats, err := s.repo.GetUserChats(userID)
	if err != nil {
		return nil, err
	}

	var response []ChatResponse
	for _, chat := range chats {
		chatResp, err := s.getChatResponse(chat.ID, userID)
		if err != nil {
			continue
		}
		response = append(response, *chatResp)
	}
	return response, nil
}

func (s *Service) GetChatByID(chatID, userID int) (*ChatResponse, error) {
	// Проверяем, является ли пользователь участником
	isMember, err := s.repo.IsMember(chatID, userID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("вы не являетесь участником чата")
	}

	return s.getChatResponse(chatID, userID)
}

func (s *Service) getChatResponse(chatID, userID int) (*ChatResponse, error) {
	chat, err := s.repo.GetChatByID(chatID)
	if err != nil {
		return nil, err
	}

	members, err := s.repo.GetChatMembers(chatID)
	if err != nil {
		return nil, err
	}

	lastMessage, _ := s.repo.GetLastMessage(chatID)
	unreadCount, _ := s.repo.GetUnreadCount(chatID, userID)

	memberResponses := make([]ChatMemberResponse, 0)
	for _, m := range members {
		var userName string
		s.repo.db.QueryRow("SELECT name FROM users WHERE id = $1", m.UserID).Scan(&userName)
		memberResponses = append(memberResponses, ChatMemberResponse{
			UserID:   m.UserID,
			UserName: userName,
			JoinedAt: m.JoinedAt,
		})
	}

	response := &ChatResponse{
		ID:          chat.ID,
		Name:        chat.Name,
		Type:        chat.Type,
		Members:     memberResponses,
		UnreadCount: unreadCount,
		CreatedAt:   chat.CreatedAt,
	}

	if lastMessage != nil {
		response.LastMessage = lastMessage.Text
		response.LastMessageAt = lastMessage.CreatedAt
	}

	// Для приватных чатов генерируем имя из имен участников
	if chat.Type == "private" && chat.Name == "" {
		// Берем имя другого участника
		for _, m := range memberResponses {
			if m.UserID != userID {
				response.Name = m.UserName
				break
			}
		}
	}

	return response, nil
}

func (s *Service) SendMessage(chatID, userID int, req *SendMessageRequest) (*MessageResponse, error) {
	// Проверяем, является ли пользователь участником
	isMember, err := s.repo.IsMember(chatID, userID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("вы не являетесь участником чата")
	}

	msg := &models.Message{
		ChatID:  chatID,
		UserID:  userID,
		Text:    req.Text,
		FileURL: req.FileURL,
	}

	if err := s.repo.SendMessage(msg); err != nil {
		return nil, err
	}

	// Получаем имя пользователя
	var userName, userPhoto string
	s.repo.db.QueryRow("SELECT name, photo FROM users WHERE id = $1", userID).Scan(&userName, &userPhoto)

	return &MessageResponse{
		ID:        msg.ID,
		UserID:    msg.UserID,
		UserName:  userName,
		UserPhoto: userPhoto,
		Text:      msg.Text,
		FileURL:   msg.FileURL,
		IsRead:    msg.IsRead,
		IsMine:    true,
		CreatedAt: msg.CreatedAt,
	}, nil
}

func (s *Service) GetMessages(chatID, userID int, query *GetMessagesQuery) ([]MessageResponse, error) {
	// Проверяем, является ли пользователь участником
	isMember, err := s.repo.IsMember(chatID, userID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("вы не являетесь участником чата")
	}

	limit := query.Limit
	if limit == 0 {
		limit = 50
	}

	messages, err := s.repo.GetMessages(chatID, limit, query.Offset)
	if err != nil {
		return nil, err
	}

	// Отмечаем сообщения как прочитанные
	s.repo.MarkChatAsRead(chatID, userID)

	var response []MessageResponse
	for _, msg := range messages {
		var userName, userPhoto string
		s.repo.db.QueryRow("SELECT name, photo FROM users WHERE id = $1", msg.UserID).Scan(&userName, &userPhoto)

		response = append(response, MessageResponse{
			ID:        msg.ID,
			UserID:    msg.UserID,
			UserName:  userName,
			UserPhoto: userPhoto,
			Text:      msg.Text,
			FileURL:   msg.FileURL,
			IsRead:    msg.IsRead,
			IsMine:    msg.UserID == userID,
			CreatedAt: msg.CreatedAt,
		})
	}
	return response, nil
}

func (s *Service) MarkMessageAsRead(messageID, userID int) error {
	return s.repo.MarkMessageAsRead(messageID, userID)
}

func (s *Service) DeleteMessage(messageID, userID int) error {
	return s.repo.DeleteMessage(messageID, userID)
}

func (s *Service) LeaveChat(chatID, userID int) error {
	return s.repo.LeaveChat(chatID, userID)
}
