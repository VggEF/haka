package notifications

import (
	"database/sql"
	"fmt"
	"student-app/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// ========== Уведомления ==========
func (r *Repository) Create(notification *models.Notification) error {
	query := `
        INSERT INTO notifications (user_id, title, message, type, data)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at
    `
	return r.db.QueryRow(query, notification.UserID, notification.Title,
		notification.Message, notification.Type, notification.Data).Scan(&notification.ID, &notification.CreatedAt)
}

func (r *Repository) GetUserNotifications(userID int, notificationType string, isRead *bool, limit, offset int) ([]models.Notification, error) {
	if limit == 0 {
		limit = 50
	}

	query := `
        SELECT id, user_id, title, message, type, is_read, data, created_at
        FROM notifications
        WHERE user_id = $1
    `
	args := []interface{}{userID}
	i := 2

	if notificationType != "" {
		query += fmt.Sprintf(" AND type = $%d", i)
		args = append(args, notificationType)
		i++
	}
	if isRead != nil {
		query += fmt.Sprintf(" AND is_read = $%d", i)
		args = append(args, *isRead)
		i++
	}

	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var n models.Notification
		err := rows.Scan(&n.ID, &n.UserID, &n.Title, &n.Message,
			&n.Type, &n.IsRead, &n.Data, &n.CreatedAt)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, n)
	}
	return notifications, nil
}

func (r *Repository) GetUnreadCount(userID int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM notifications WHERE user_id = $1 AND is_read = false`
	err := r.db.QueryRow(query, userID).Scan(&count)
	return count, err
}

func (r *Repository) MarkAsRead(notificationID, userID int) error {
	query := `UPDATE notifications SET is_read = true WHERE id = $1 AND user_id = $2`
	_, err := r.db.Exec(query, notificationID, userID)
	return err
}

func (r *Repository) MarkAllAsRead(userID int) error {
	query := `UPDATE notifications SET is_read = true WHERE user_id = $1 AND is_read = false`
	_, err := r.db.Exec(query, userID)
	return err
}

func (r *Repository) DeleteNotification(notificationID, userID int) error {
	query := `DELETE FROM notifications WHERE id = $1 AND user_id = $2`
	_, err := r.db.Exec(query, notificationID, userID)
	return err
}

func (r *Repository) DeleteOldNotifications(days int) error {
	query := `DELETE FROM notifications WHERE created_at < NOW() - INTERVAL '1 day' * $1`
	_, err := r.db.Exec(query, days)
	return err
}

// ========== Массовые уведомления ==========
func (r *Repository) SendToAllUsers(title, message, notificationType, data string) error {
	query := `
        INSERT INTO notifications (user_id, title, message, type, data)
        SELECT id, $1, $2, $3, $4 FROM users
    `
	_, err := r.db.Exec(query, title, message, notificationType, data)
	return err
}

func (r *Repository) SendToGroup(groupName, title, message, notificationType, data string) error {
	query := `
        INSERT INTO notifications (user_id, title, message, type, data)
        SELECT id, $1, $2, $3, $4 FROM users WHERE group_name = $5
    `
	_, err := r.db.Exec(query, title, message, notificationType, data, groupName)
	return err
}

func (r *Repository) SendToRole(role, title, message, notificationType, data string) error {
	query := `
        INSERT INTO notifications (user_id, title, message, type, data)
        SELECT id, $1, $2, $3, $4 FROM users WHERE role = $5
    `
	_, err := r.db.Exec(query, title, message, notificationType, data, role)
	return err
}
