package library

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

func (r *Repository) Create(book *models.Book) error {
	query := `
        INSERT INTO books (title, author, publisher, year, isbn, description, cover_url, file_path, category, tags, available_copies)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
        RETURNING id
    `
	return r.db.QueryRow(query, book.Title, book.Author, book.Publisher,
		book.Year, book.ISBN, book.Description, book.CoverURL, book.FilePath,
		book.Category, book.Tags, book.AvailableCopies).Scan(&book.ID)
}

func (r *Repository) Update(id int, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	setParts := make([]string, 0, len(updates))
	args := make([]interface{}, 0, len(updates)+1)
	i := 1

	for key, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE books SET %s WHERE id = $%d", strings.Join(setParts, ", "), i)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Repository) Delete(id int) error {
	query := `DELETE FROM books WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) FindByID(id int) (*models.Book, error) {
	var book models.Book
	query := `
        SELECT id, title, author, publisher, year, isbn, description, cover_url, file_path, category, tags, available_copies
        FROM books
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author,
		&book.Publisher, &book.Year, &book.ISBN, &book.Description, &book.CoverURL,
		&book.FilePath, &book.Category, &book.Tags, &book.AvailableCopies)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *Repository) Search(query string, author string, year string, limit, offset int) ([]models.Book, error) {
	sqlQuery := `
        SELECT id, title, author, publisher, year, description, cover_url, category
        FROM books
        WHERE 1=1
    `
	args := []interface{}{}
	i := 1

	if query != "" {
		sqlQuery += fmt.Sprintf(" AND (title ILIKE $%d OR author ILIKE $%d)", i, i)
		args = append(args, "%"+query+"%")
		i++
	}
	if author != "" {
		sqlQuery += fmt.Sprintf(" AND author ILIKE $%d", i)
		args = append(args, "%"+author+"%")
		i++
	}
	if year != "" {
		sqlQuery += fmt.Sprintf(" AND year = $%d", i)
		args = append(args, year)
		i++
	}

	sqlQuery += fmt.Sprintf(" ORDER BY title LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Publisher,
			&book.Year, &book.Description, &book.CoverURL, &book.Category)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *Repository) GetAll(limit, offset int) ([]models.Book, error) {
	query := `
        SELECT id, title, author, publisher, year, description, cover_url, category
        FROM books
        ORDER BY title
        LIMIT $1 OFFSET $2
    `
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Publisher,
			&book.Year, &book.Description, &book.CoverURL, &book.Category)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *Repository) SaveCache(searchQuery string, data string) error {
	query := `
        INSERT INTO library_cache (search_query, data)
        VALUES ($1, $2)
        ON CONFLICT (search_query) DO UPDATE SET data = EXCLUDED.data, fetched_at = NOW()
    `
	_, err := r.db.Exec(query, searchQuery, data)
	return err
}

func (r *Repository) GetCache(searchQuery string) (string, error) {
	var data string
	query := `SELECT data FROM library_cache WHERE search_query = $1 AND fetched_at > NOW() - INTERVAL '1 hour'`
	err := r.db.QueryRow(query, searchQuery).Scan(&data)
	if err != nil {
		return "", err
	}
	return data, nil
}
