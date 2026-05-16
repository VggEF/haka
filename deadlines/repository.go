package deadlines

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

func (r *Repository) Create(deadline *models.Deadline) error {
	query := `
        INSERT INTO deadlines (title, subject, group_name, due_date, due_time, priority, description, created_by, assigned_to)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id, created_at
    `
	return r.db.QueryRow(query, deadline.Title, deadline.Subject, deadline.GroupName,
		deadline.DueDate, deadline.DueTime, deadline.Priority, deadline.Description,
		deadline.CreatedBy, deadline.AssignedTo).Scan(&deadline.ID, &deadline.CreatedAt)
}

func (r *Repository) Update(id int, updates map[string]interface{}) error {
	setParts := make([]string, 0, len(updates))
	args := make([]interface{}, 0, len(updates)+1)
	i := 1

	for key, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE deadlines SET %s WHERE id = $%d", strings.Join(setParts, ", "), i)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Repository) Delete(id int) error {
	query := `DELETE FROM deadlines WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) FindByID(id int) (*models.Deadline, error) {
	var deadline models.Deadline
	query := `
        SELECT id, title, subject, group_name, due_date, due_time, priority, description, status, created_by, assigned_to, created_at
        FROM deadlines
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(
		&deadline.ID, &deadline.Title, &deadline.Subject, &deadline.GroupName,
		&deadline.DueDate, &deadline.DueTime, &deadline.Priority, &deadline.Description,
		&deadline.Status, &deadline.CreatedBy, &deadline.AssignedTo, &deadline.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &deadline, nil
}

func (r *Repository) FindByUser(userID int, query *GetDeadlinesQuery) ([]models.Deadline, error) {
	sqlQuery := `
        SELECT id, title, subject, group_name, due_date, due_time, priority, description, status, created_by, created_at
        FROM deadlines
        WHERE assigned_to = $1 OR group_name IN (SELECT group_name FROM users WHERE id = $1)
    `
	args := []interface{}{userID}
	i := 2

	if query.Subject != "" {
		sqlQuery += fmt.Sprintf(" AND subject ILIKE $%d", i)
		args = append(args, "%"+query.Subject+"%")
		i++
	}
	if query.Status != "" {
		sqlQuery += fmt.Sprintf(" AND status = $%d", i)
		args = append(args, query.Status)
		i++
	}
	if query.Priority != "" {
		sqlQuery += fmt.Sprintf(" AND priority = $%d", i)
		args = append(args, query.Priority)
		i++
	}
	if query.DueFrom != "" {
		sqlQuery += fmt.Sprintf(" AND due_date >= $%d", i)
		args = append(args, query.DueFrom)
		i++
	}
	if query.DueTo != "" {
		sqlQuery += fmt.Sprintf(" AND due_date <= $%d", i)
		args = append(args, query.DueTo)
		i++
	}

	limit := query.Limit
	if limit == 0 {
		limit = 50
	}
	sqlQuery += fmt.Sprintf(" ORDER BY due_date ASC LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, query.Offset)

	rows, err := r.db.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deadlines []models.Deadline
	for rows.Next() {
		var d models.Deadline
		err := rows.Scan(&d.ID, &d.Title, &d.Subject, &d.GroupName,
			&d.DueDate, &d.DueTime, &d.Priority, &d.Description,
			&d.Status, &d.CreatedBy, &d.CreatedAt)
		if err != nil {
			return nil, err
		}
		deadlines = append(deadlines, d)
	}
	return deadlines, nil
}

func (r *Repository) FindByGroup(groupName string, query *GetDeadlinesQuery) ([]models.Deadline, error) {
	sqlQuery := `
        SELECT id, title, subject, group_name, due_date, due_time, priority, description, status, created_by, created_at
        FROM deadlines
        WHERE group_name = $1
    `
	args := []interface{}{groupName}
	i := 2

	if query.Subject != "" {
		sqlQuery += fmt.Sprintf(" AND subject ILIKE $%d", i)
		args = append(args, "%"+query.Subject+"%")
		i++
	}
	if query.Status != "" {
		sqlQuery += fmt.Sprintf(" AND status = $%d", i)
		args = append(args, query.Status)
		i++
	}
	if query.Priority != "" {
		sqlQuery += fmt.Sprintf(" AND priority = $%d", i)
		args = append(args, query.Priority)
		i++
	}

	limit := query.Limit
	if limit == 0 {
		limit = 50
	}
	sqlQuery += fmt.Sprintf(" ORDER BY due_date ASC LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, query.Offset)

	rows, err := r.db.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deadlines []models.Deadline
	for rows.Next() {
		var d models.Deadline
		err := rows.Scan(&d.ID, &d.Title, &d.Subject, &d.GroupName,
			&d.DueDate, &d.DueTime, &d.Priority, &d.Description,
			&d.Status, &d.CreatedBy, &d.CreatedAt)
		if err != nil {
			return nil, err
		}
		deadlines = append(deadlines, d)
	}
	return deadlines, nil
}

func (r *Repository) UpdateStatus(id int, status string) error {
	query := `UPDATE deadlines SET status = $1 WHERE id = $2`
	_, err := r.db.Exec(query, status, id)
	return err
}

func (r *Repository) GetUpcomingDeadlines(userID int, limit int) ([]models.Deadline, error) {
	query := `
        SELECT id, title, subject, due_date, due_time, priority
        FROM deadlines
        WHERE (assigned_to = $1 OR group_name IN (SELECT group_name FROM users WHERE id = $1))
            AND status = 'pending'
            AND due_date >= CURRENT_DATE
        ORDER BY due_date ASC
        LIMIT $2
    `
	rows, err := r.db.Query(query, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deadlines []models.Deadline
	for rows.Next() {
		var d models.Deadline
		err := rows.Scan(&d.ID, &d.Title, &d.Subject, &d.DueDate, &d.DueTime, &d.Priority)
		if err != nil {
			return nil, err
		}
		deadlines = append(deadlines, d)
	}
	return deadlines, nil
}
