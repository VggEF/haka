package library

import (
	"encoding/json"
	"student-app/internal/models"
)

type Service struct {
	repo      *Repository
	apiClient *APIClient
}

func NewService(repo *Repository, apiClient *APIClient) *Service {
	return &Service{
		repo:      repo,
		apiClient: apiClient,
	}
}

func (s *Service) AddBook(req *AddBookRequest) (*models.Book, error) {
	book := &models.Book{
		Title:           req.Title,
		Author:          req.Author,
		Publisher:       req.Publisher,
		Year:            req.Year,
		ISBN:            req.ISBN,
		Description:     req.Description,
		CoverURL:        req.CoverURL,
		FilePath:        req.FilePath,
		Category:        req.Category,
		Tags:            req.Tags,
		AvailableCopies: 1,
	}

	if err := s.repo.Create(book); err != nil {
		return nil, err
	}
	return book, nil
}

func (s *Service) UpdateBook(id int, req *UpdateBookRequest) error {
	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Author != "" {
		updates["author"] = req.Author
	}
	if req.Publisher != "" {
		updates["publisher"] = req.Publisher
	}
	if req.Year != "" {
		updates["year"] = req.Year
	}
	if req.ISBN != "" {
		updates["isbn"] = req.ISBN
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.CoverURL != "" {
		updates["cover_url"] = req.CoverURL
	}
	if req.FilePath != "" {
		updates["file_path"] = req.FilePath
	}
	if req.Category != "" {
		updates["category"] = req.Category
	}
	if len(req.Tags) > 0 {
		updates["tags"] = req.Tags
	}

	return s.repo.Update(id, updates)
}

func (s *Service) DeleteBook(id int) error {
	return s.repo.Delete(id)
}

func (s *Service) GetBookByID(id int) (*BookResponse, error) {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		Author:          book.Author,
		Publisher:       book.Publisher,
		Year:            book.Year,
		ISBN:            book.ISBN,
		Description:     book.Description,
		CoverURL:        book.CoverURL,
		FilePath:        book.FilePath,
		Category:        book.Category,
		Tags:            book.Tags,
		AvailableCopies: book.AvailableCopies,
	}, nil
}

func (s *Service) SearchBooks(query *SearchBooksQuery) ([]BookResponse, error) {
	var books []models.Book
	var err error

	if query.Query != "" {
		// Сначала проверяем кэш
		cached, _ := s.repo.GetCache(query.Query)
		if cached != "" {
			var cachedBooks []models.Book
			json.Unmarshal([]byte(cached), &cachedBooks)
			books = cachedBooks
		} else {
			// Ищем во внешнем API
			externalBooks, err := s.apiClient.SearchBooks(query.Query, query.Author, query.Year)
			if err == nil && len(externalBooks) > 0 {
				for _, eb := range externalBooks {
					books = append(books, models.Book{
						Title:       eb.Title,
						Author:      eb.Author,
						Publisher:   eb.Publisher,
						Year:        eb.Year,
						Description: eb.Description,
						CoverURL:    eb.CoverURL,
					})
				}
				// Сохраняем в кэш
				data, _ := json.Marshal(books)
				s.repo.SaveCache(query.Query, string(data))
			}
		}
	}

	if len(books) == 0 {
		limit := query.Limit
		if limit == 0 {
			limit = 20
		}
		books, err = s.repo.Search(query.Query, query.Author, query.Year, limit, query.Offset)
		if err != nil {
			return nil, err
		}
	}

	var response []BookResponse
	for _, book := range books {
		response = append(response, BookResponse{
			ID:              book.ID,
			Title:           book.Title,
			Author:          book.Author,
			Publisher:       book.Publisher,
			Year:            book.Year,
			ISBN:            book.ISBN,
			Description:     book.Description,
			CoverURL:        book.CoverURL,
			FilePath:        book.FilePath,
			Category:        book.Category,
			Tags:            book.Tags,
			AvailableCopies: book.AvailableCopies,
		})
	}
	return response, nil
}

func (s *Service) GetAllBooks(limit, offset int) ([]BookResponse, error) {
	books, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	var response []BookResponse
	for _, book := range books {
		response = append(response, BookResponse{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Publisher:   book.Publisher,
			Year:        book.Year,
			Description: book.Description,
			CoverURL:    book.CoverURL,
			Category:    book.Category,
		})
	}
	return response, nil
}
