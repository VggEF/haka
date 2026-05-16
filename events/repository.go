package events

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"student-app/internal/models"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(event *models.Event) error {
	query := `
        INSERT INTO events (title, short_text, full_text, date, time, type, category,
                           location, price, organizer, image_url, available_spots, contact, created_by)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
        RETURNING id, created_at
    `
	return r.db.QueryRow(query, event.Title, event.ShortText, event.FullText,
		event.Date, event.Time, event.Type, event.Category, event.Location,
		event.Price, event.Organizer, event.ImageURL, event.AvailableSpots,
		event.Contact, event.CreatedBy).Scan(&event.ID, &event.CreatedAt)
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
	query := fmt.Sprintf("UPDATE events SET %s WHERE id = $%d", strings.Join(setParts, ", "), i)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Repository) Delete(id int) error {
	query := `DELETE FROM events WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) FindByID(id int) (*models.Event, error) {
	var event models.Event
	query := `
        SELECT id, title, short_text, full_text, date, time, type, category,
               location, price, organizer, image_url, available_spots, contact,
               registrations, created_by, created_at
        FROM events
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(
		&event.ID, &event.Title, &event.ShortText, &event.FullText,
		&event.Date, &event.Time, &event.Type, &event.Category,
		&event.Location, &event.Price, &event.Organizer, &event.ImageURL,
		&event.AvailableSpots, &event.Contact, &event.Registrations,
		&event.CreatedBy, &event.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *Repository) GetAll(eventType, category, dateFrom, dateTo string, limit, offset int) ([]models.Event, error) {
	query := `
        SELECT id, title, short_text, date, time, type, category, location, price, organizer, image_url, available_spots, registrations
        FROM events
        WHERE 1=1
    `
	args := []interface{}{}
	i := 1

	if eventType != "" {
		query += fmt.Sprintf(" AND type = $%d", i)
		args = append(args, eventType)
		i++
	}
	if category != "" {
		query += fmt.Sprintf(" AND category = $%d", i)
		args = append(args, category)
		i++
	}
	if dateFrom != "" {
		query += fmt.Sprintf(" AND date >= $%d", i)
		args = append(args, dateFrom)
		i++
	}
	if dateTo != "" {
		query += fmt.Sprintf(" AND date <= $%d", i)
		args = append(args, dateTo)
		i++
	}

	query += fmt.Sprintf(" ORDER BY date ASC LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.Title, &event.ShortText, &event.Date,
			&event.Time, &event.Type, &event.Category, &event.Location,
			&event.Price, &event.Organizer, &event.ImageURL, &event.AvailableSpots,
			&event.Registrations)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *Repository) AddRegistration(eventID, userID int) error {
	// Сначала получаем текущие регистрации
	var registrationsJSON string
	query := `SELECT registrations FROM events WHERE id = $1`
	err := r.db.QueryRow(query, eventID).Scan(&registrationsJSON)
	if err != nil {
		return err
	}

	var registrations []map[string]interface{}
	if registrationsJSON != "" {
		json.Unmarshal([]byte(registrationsJSON), &registrations)
	}

	// Проверяем, не зарегистрирован ли уже
	for _, reg := range registrations {
		if uid, ok := reg["user_id"]; ok && int(uid.(float64)) == userID {
			return fmt.Errorf("user already registered")
		}
	}

	// Добавляем новую регистрацию
	registrations = append(registrations, map[string]interface{}{
		"user_id":       userID,
		"registered_at": time.Now(),
	})

	newRegistrationsJSON, _ := json.Marshal(registrations)
	updateQuery := `UPDATE events SET registrations = $1 WHERE id = $2`
	_, err = r.db.Exec(updateQuery, string(newRegistrationsJSON), eventID)
	return err
}

func (r *Repository) RemoveRegistration(eventID, userID int) error {
	var registrationsJSON string
	query := `SELECT registrations FROM events WHERE id = $1`
	err := r.db.QueryRow(query, eventID).Scan(&registrationsJSON)
	if err != nil {
		return err
	}

	var registrations []map[string]interface{}
	if registrationsJSON != "" {
		json.Unmarshal([]byte(registrationsJSON), &registrations)
	}

	// Удаляем регистрацию
	newRegistrations := []map[string]interface{}{}
	for _, reg := range registrations {
		if uid, ok := reg["user_id"]; ok && int(uid.(float64)) != userID {
			newRegistrations = append(newRegistrations, reg)
		}
	}

	newRegistrationsJSON, _ := json.Marshal(newRegistrations)
	updateQuery := `UPDATE events SET registrations = $1 WHERE id = $2`
	_, err = r.db.Exec(updateQuery, string(newRegistrationsJSON), eventID)
	return err
}

func (r *Repository) IsRegistered(eventID, userID int) (bool, error) {
	var registrationsJSON string
	query := `SELECT registrations FROM events WHERE id = $1`
	err := r.db.QueryRow(query, eventID).Scan(&registrationsJSON)
	if err != nil {
		return false, err
	}

	var registrations []map[string]interface{}
	if registrationsJSON != "" {
		json.Unmarshal([]byte(registrationsJSON), &registrations)
	}

	for _, reg := range registrations {
		if uid, ok := reg["user_id"]; ok && int(uid.(float64)) == userID {
			return true, nil
		}
	}
	return false, nil
}

func (r *Repository) GetRegistrationsCount(eventID int) (int, error) {
	var registrationsJSON string
	query := `SELECT registrations FROM events WHERE id = $1`
	err := r.db.QueryRow(query, eventID).Scan(&registrationsJSON)
	if err != nil {
		return 0, err
	}

	var registrations []map[string]interface{}
	if registrationsJSON != "" {
		json.Unmarshal([]byte(registrationsJSON), &registrations)
	}
	return len(registrations), nil
}
