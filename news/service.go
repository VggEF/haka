package news

import (
	"student-app/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(req *CreateNewsRequest, userID int) (*models.News, error) {
	news := &models.News{
		Title:     req.Title,
		ShortText: req.ShortText,
		FullText:  req.FullText,
		ImageURL:  req.ImageURL,
		Category:  req.Category,
		IsPinned:  req.IsPinned,
		CreatedBy: userID,
	}

	if err := s.repo.Create(news); err != nil {
		return nil, err
	}
	return news, nil
}

func (s *Service) Update(id int, req *UpdateNewsRequest) error {
	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.ShortText != "" {
		updates["short_text"] = req.ShortText
	}
	if req.FullText != "" {
		updates["full_text"] = req.FullText
	}
	if req.ImageURL != "" {
		updates["image_url"] = req.ImageURL
	}
	if req.Category != "" {
		updates["category"] = req.Category
	}
	updates["is_pinned"] = req.IsPinned

	return s.repo.Update(id, updates)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *Service) GetByID(id int) (*NewsResponse, error) {
	news, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Увеличиваем счетчик просмотров
	s.repo.IncrementViews(id)

	return &NewsResponse{
		ID:        news.ID,
		Title:     news.Title,
		ShortText: news.ShortText,
		FullText:  news.FullText,
		ImageURL:  news.ImageURL,
		Category:  news.Category,
		Date:      news.Date,
		IsPinned:  news.IsPinned,
		Views:     news.Views,
		CreatedAt: news.CreatedAt,
	}, nil
}

func (s *Service) GetAll(query *GetNewsQuery) ([]NewsResponse, int, error) {
	limit := query.Limit
	if limit == 0 {
		limit = 20
	}

	newsList, err := s.repo.GetAll(query.Category, query.Pinned, limit, query.Offset)
	if err != nil {
		return nil, 0, err
	}

	var response []NewsResponse
	for _, news := range newsList {
		response = append(response, NewsResponse{
			ID:        news.ID,
			Title:     news.Title,
			ShortText: news.ShortText,
			FullText:  news.FullText,
			ImageURL:  news.ImageURL,
			Category:  news.Category,
			Date:      news.Date,
			IsPinned:  news.IsPinned,
			Views:     news.Views,
		})
	}

	return response, len(newsList), nil
}

func (s *Service) GetPinned() ([]NewsResponse, error) {
	newsList, err := s.repo.GetPinned()
	if err != nil {
		return nil, err
	}

	var response []NewsResponse
	for _, news := range newsList {
		response = append(response, NewsResponse{
			ID:        news.ID,
			Title:     news.Title,
			ShortText: news.ShortText,
			ImageURL:  news.ImageURL,
			Category:  news.Category,
			Date:      news.Date,
			Views:     news.Views,
		})
	}
	return response, nil
}
