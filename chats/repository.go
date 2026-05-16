package chats

import (
	"database/sql"
	"student-app/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// ========== Чаты ==========
func (r *Repository) CreateChat(chat *models.Chat) error {
	query := `
        INSERT INTO chats (name, type)
        VALUES ($1, $2)
        RETURNING id, created_at
    `
	return r.db.QueryRow(query, chat.Name, chat.Type).Scan(&chat.ID, &chat.CreatedAt)
}

func (r *Repository) AddMembers(chatID int, userIDs []int) error {
	query := `INSERT INTO chat_members (chat_id, user_id) VALUES ($1, $2)`
	for _, userID := range userIDs {
		_, err := r.db.Exec(query, chatID, userID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) GetChatByID(chatID int) (*models.Chat, error) {
	var chat models.Chat
	query := `SELECT id, name, type, created_at FROM chats WHERE id = $1`
	err := r.db.QueryRow(query, chatID).Scan(&chat.ID, &chat.Name, &chat.Type, &chat.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func (r *Repository) GetUserChats(userID int) ([]models.Chat, error) {
	query := `
        SELECT c.id, c.name, c.type, c.created_at
        FROM chats c
        JOIN chat_members cm ON c.id = cm.chat_id
        WHERE cm.user_id = $1
        ORDER BY c.created_at DESC
    `
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []models.Chat
	for rows.Next() {
		var chat models.Chat
		err := rows.Scan(&chat.ID, &chat.Name, &chat.Type, &chat.CreatedAt)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}

func (r *Repository) GetChatMembers(chatID int) ([]models.ChatMember, error) {
	query := `
        SELECT cm.user_id, cm.joined_at, u.name
        FROM chat_members cm
        JOIN users u ON cm.user_id = u.id
        WHERE cm.chat_id = $1
    `
	rows, err := r.db.Query(query, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []models.ChatMember
	for rows.Next() {
		var member models.ChatMember
		var userName string
		err := rows.Scan(&member.UserID, &member.JoinedAt, &userName)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}

func (r *Repository) IsMember(chatID, userID int) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM chat_members WHERE chat_id = $1 AND user_id = $2)`
	err := r.db.QueryRow(query, chatID, userID).Scan(&exists)
	return exists, err
}

func (r *Repository) LeaveChat(chatID, userID int) error {
	query := `DELETE FROM chat_members WHERE chat_id = $1 AND user_id = $2`
	_, err := r.db.Exec(query, chatID, userID)
	return err
}

// ========== Сообщения ==========
func (r *Repository) SendMessage(msg *models.Message) error {
	query := `
        INSERT INTO messages (chat_id, user_id, text, file_url)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `
	return r.db.QueryRow(query, msg.ChatID, msg.UserID, msg.Text, msg.FileURL).Scan(&msg.ID, &msg.CreatedAt)
}

func (r *Repository) GetMessages(chatID int, limit, offset int) ([]models.Message, error) {
	if limit == 0 {
		limit = 50
	}
	query := `
        SELECT id, chat_id, user_id, text, file_url, is_read, created_at
        FROM messages
        WHERE chat_id = $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `
	rows, err := r.db.Query(query, chatID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.ChatID, &msg.UserID, &msg.Text,
			&msg.FileURL, &msg.IsRead, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	// Переворачиваем для хронологического порядка
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
	return messages, nil
}

func (r *Repository) MarkMessageAsRead(messageID, userID int) error {
	query := `
        UPDATE messages
        SET is_read = true
        WHERE id = $1 AND user_id != $2
    `
	_, err := r.db.Exec(query, messageID, userID)
	return err
}

func (r *Repository) MarkChatAsRead(chatID, userID int) error {
	query := `
        UPDATE messages
        SET is_read = true
        WHERE chat_id = $1 AND user_id != $2 AND is_read = false
    `
	_, err := r.db.Exec(query, chatID, userID)
	return err
}

func (r *Repository) GetUnreadCount(chatID, userID int) (int, error) {
	var count int
	query := `
        SELECT COUNT(*)
        FROM messages
        WHERE chat_id = $1 AND user_id != $2 AND is_read = false
    `
	err := r.db.QueryRow(query, chatID, userID).Scan(&count)
	return count, err
}

func (r *Repository) GetLastMessage(chatID int) (*models.Message, error) {
	var msg models.Message
	query := `
        SELECT id, user_id, text, created_at
        FROM messages
        WHERE chat_id = $1
        ORDER BY created_at DESC
        LIMIT 1
    `
	err := r.db.QueryRow(query, chatID).Scan(&msg.ID, &msg.UserID, &msg.Text, &msg.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func (r *Repository) DeleteMessage(messageID, userID int) error {
	// Только автор может удалить
	query := `DELETE FROM messages WHERE id = $1 AND user_id = $2`
	_, err := r.db.Exec(query, messageID, userID)
	return err
}
