package news

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

func (r *Repository) Create(news *models.News) error {
	query := `
        INSERT INTO news (title, short_text, full_text, image_url, category, is_pinned, created_by)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, created_at
    `
	return r.db.QueryRow(query, news.Title, news.ShortText, news.FullText,
		news.ImageURL, news.Category, news.IsPinned, news.CreatedBy).Scan(&news.ID, &news.CreatedAt)
}

func (r *Repository) Update(id int, updates map[string]interface{}) error {
	query := `UPDATE news SET updated_at = NOW()`
	args := []interface{}{}
	i := 1

	if title, ok := updates["title"]; ok {
		query += fmt.Sprintf(", title = $%d", i)
		args = append(args, title)
		i++
	}
	if shortText, ok := updates["short_text"]; ok {
		query += fmt.Sprintf(", short_text = $%d", i)
		args = append(args, shortText)
		i++
	}
	if fullText, ok := updates["full_text"]; ok {
		query += fmt.Sprintf(", full_text = $%d", i)
		args = append(args, fullText)
		i++
	}
	if imageURL, ok := updates["image_url"]; ok {
		query += fmt.Sprintf(", image_url = $%d", i)
		args = append(args, imageURL)
		i++
	}
	if category, ok := updates["category"]; ok {
		query += fmt.Sprintf(", category = $%d", i)
		args = append(args, category)
		i++
	}
	if isPinned, ok := updates["is_pinned"]; ok {
		query += fmt.Sprintf(", is_pinned = $%d", i)
		args = append(args, isPinned)
		i++
	}

	query += fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Repository) Delete(id int) error {
	query := `DELETE FROM news WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) FindByID(id int) (*models.News, error) {
	var news models.News
	query := `
        SELECT id, title, short_text, full_text, image_url, category, date, is_pinned, views, created_by, created_at
        FROM news
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(
		&news.ID, &news.Title, &news.ShortText, &news.FullText, &news.ImageURL,
		&news.Category, &news.Date, &news.IsPinned, &news.Views, &news.CreatedBy, &news.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *Repository) IncrementViews(id int) error {
	query := `UPDATE news SET views = views + 1 WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) GetAll(category string, pinned bool, limit, offset int) ([]models.News, error) {
	query := `
        SELECT id, title, short_text, full_text, image_url, category, date, is_pinned, views
        FROM news
        WHERE 1=1
    `
	args := []interface{}{}
	i := 1

	if category != "" {
		query += fmt.Sprintf(" AND category = $%d", i)
		args = append(args, category)
		i++
	}
	if pinned {
		query += " AND is_pinned = true"
	}

	query += fmt.Sprintf(" ORDER BY is_pinned DESC, date DESC LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsList []models.News
	for rows.Next() {
		var news models.News
		err := rows.Scan(&news.ID, &news.Title, &news.ShortText, &news.FullText,
			&news.ImageURL, &news.Category, &news.Date, &news.IsPinned, &news.Views)
		if err != nil {
			return nil, err
		}
		newsList = append(newsList, news)
	}
	return newsList, nil
}

func (r *Repository) GetPinned() ([]models.News, error) {
	query := `
        SELECT id, title, short_text, image_url, category, date, views
        FROM news
        WHERE is_pinned = true
        ORDER BY date DESC
        LIMIT 3
    `
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsList []models.News
	for rows.Next() {
		var news models.News
		err := rows.Scan(&news.ID, &news.Title, &news.ShortText, &news.ImageURL,
			&news.Category, &news.Date, &news.Views)
		if err != nil {
			return nil, err
		}
		newsList = append(newsList, news)
	}
	return newsList, nil
}
