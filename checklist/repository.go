package checklist

import (
	"database/sql"
	"fmt"
	"strings"
	"student-app/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// ========== Элементы чек-листа ==========
func (r *Repository) CreateItem(item *models.ChecklistItem) error {
	query := `
        INSERT INTO checklist_items (title, description, category, xp, sort_order)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	return r.db.QueryRow(query, item.Title, item.Description, item.Category,
		item.XP, item.SortOrder).Scan(&item.ID)
}

func (r *Repository) UpdateItem(id int, updates map[string]interface{}) error {
	setParts := make([]string, 0, len(updates))
	args := make([]interface{}, 0, len(updates)+1)
	i := 1

	for key, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE checklist_items SET %s WHERE id = $%d", strings.Join(setParts, ", "), i)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Repository) DeleteItem(id int) error {
	query := `DELETE FROM checklist_items WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) FindItemByID(id int) (*models.ChecklistItem, error) {
	var item models.ChecklistItem
	query := `
        SELECT id, title, description, category, xp, sort_order, is_active
        FROM checklist_items
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(&item.ID, &item.Title, &item.Description,
		&item.Category, &item.XP, &item.SortOrder, &item.IsActive)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *Repository) GetAllItems(category string, isActive bool) ([]models.ChecklistItem, error) {
	query := `
        SELECT id, title, description, category, xp, sort_order, is_active
        FROM checklist_items
        WHERE 1=1
    `
	args := []interface{}{}
	i := 1

	if category != "" {
		query += fmt.Sprintf(" AND category = $%d", i)
		args = append(args, category)
		i++
	}
	if isActive {
		query += " AND is_active = true"
	}

	query += " ORDER BY sort_order ASC, id ASC"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.ChecklistItem
	for rows.Next() {
		var item models.ChecklistItem
		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Category,
			&item.XP, &item.SortOrder, &item.IsActive)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// ========== Прогресс пользователя ==========
func (r *Repository) CompleteItem(userID, itemID int) error {
	query := `
        INSERT INTO user_checklist (user_id, item_id, completed_at, is_completed)
        VALUES ($1, $2, NOW(), true)
        ON CONFLICT (user_id, item_id) DO UPDATE SET
            completed_at = NOW(),
            is_completed = true
    `
	_, err := r.db.Exec(query, userID, itemID)
	return err
}

func (r *Repository) IsItemCompleted(userID, itemID int) (bool, error) {
	var isCompleted bool
	query := `SELECT is_completed FROM user_checklist WHERE user_id = $1 AND item_id = $2`
	err := r.db.QueryRow(query, userID, itemID).Scan(&isCompleted)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return isCompleted, nil
}

func (r *Repository) GetUserProgress(userID int) (map[int]bool, error) {
	query := `SELECT item_id, is_completed FROM user_checklist WHERE user_id = $1`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	progress := make(map[int]bool)
	for rows.Next() {
		var itemID int
		var isCompleted bool
		if err := rows.Scan(&itemID, &isCompleted); err != nil {
			return nil, err
		}
		progress[itemID] = isCompleted
	}
	return progress, nil
}

func (r *Repository) GetUserCompletedItems(userID int) ([]models.ChecklistItem, error) {
	query := `
        SELECT ci.id, ci.title, ci.description, ci.category, ci.xp, ci.sort_order
        FROM checklist_items ci
        JOIN user_checklist uc ON ci.id = uc.item_id
        WHERE uc.user_id = $1 AND uc.is_completed = true
        ORDER BY uc.completed_at DESC
    `
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.ChecklistItem
	for rows.Next() {
		var item models.ChecklistItem
		err := rows.Scan(&item.ID, &item.Title, &item.Description,
			&item.Category, &item.XP, &item.SortOrder)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) GetUserTotalCompletedXP(userID int) (int, error) {
	query := `
        SELECT COALESCE(SUM(ci.xp), 0)
        FROM user_checklist uc
        JOIN checklist_items ci ON uc.item_id = ci.id
        WHERE uc.user_id = $1 AND uc.is_completed = true
    `
	var totalXP int
	err := r.db.QueryRow(query, userID).Scan(&totalXP)
	return totalXP, err
}
func (r *Repository) UpdateUserTotalXP(userID, xp int) error {
	query := `UPDATE users SET total_xp = total_xp + $1 WHERE id = $2`
	_, err := r.db.Exec(query, xp, userID)
	return err
}
