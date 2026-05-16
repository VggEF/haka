package notifications

import (
	"student-app/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Отправка уведомления одному пользователю
func (s *Service) SendNotification(userID int, req *SendNotificationRequest) (*models.Notification, error) {
	notificationType := req.Type
	if notificationType == "" {
		notificationType = "system"
	}

	notification := &models.Notification{
		UserID:  userID,
		Title:   req.Title,
		Message: req.Message,
		Type:    notificationType,
		Data:    req.Data,
	}

	if err := s.repo.Create(notification); err != nil {
		return nil, err
	}
	return notification, nil
}

// Отправка уведомления всем пользователям (админ)
func (s *Service) SendToAllUsers(title, message, notificationType, data string) error {
	return s.repo.SendToAllUsers(title, message, notificationType, data)
}

// Отправка уведомления группе
func (s *Service) SendToGroup(groupName, title, message, notificationType, data string) error {
	return s.repo.SendToGroup(groupName, title, message, notificationType, data)
}

// Отправка уведомления по роли
func (s *Service) SendToRole(role, title, message, notificationType, data string) error {
	return s.repo.SendToRole(role, title, message, notificationType, data)
}

// Получение уведомлений пользователя
func (s *Service) GetUserNotifications(userID int, query *GetNotificationsQuery) ([]NotificationResponse, int, error) {
	limit := query.Limit
	if limit == 0 {
		limit = 50
	}

	notifications, err := s.repo.GetUserNotifications(userID, query.Type, query.IsRead, limit, query.Offset)
	if err != nil {
		return nil, 0, err
	}

	unreadCount, _ := s.repo.GetUnreadCount(userID)

	var response []NotificationResponse
	for _, n := range notifications {
		response = append(response, NotificationResponse{
			ID:        n.ID,
			Title:     n.Title,
			Message:   n.Message,
			Type:      n.Type,
			IsRead:    n.IsRead,
			Data:      n.Data,
			CreatedAt: n.CreatedAt,
		})
	}
	return response, unreadCount, nil
}

// Получение количества непрочитанных
func (s *Service) GetUnreadCount(userID int) (int, error) {
	return s.repo.GetUnreadCount(userID)
}

// Отметить как прочитанное
func (s *Service) MarkAsRead(notificationID, userID int) error {
	return s.repo.MarkAsRead(notificationID, userID)
}

// Отметить все как прочитанные
func (s *Service) MarkAllAsRead(userID int) error {
	return s.repo.MarkAllAsRead(userID)
}

// Удалить уведомление
func (s *Service) DeleteNotification(notificationID, userID int) error {
	return s.repo.DeleteNotification(notificationID, userID)
}

// Удалить старые уведомления
func (s *Service) DeleteOldNotifications(days int) error {
	return s.repo.DeleteOldNotifications(days)
}
