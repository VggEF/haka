package events

import (
	"errors"
	"student-app/internal/models"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(req *CreateEventRequest, userID int) (*models.Event, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.New("неверный формат даты")
	}

	event := &models.Event{
		Title:          req.Title,
		ShortText:      req.ShortText,
		FullText:       req.FullText,
		Date:           date,
		Time:           req.Time,
		Type:           req.Type,
		Category:       req.Category,
		Location:       req.Location,
		Price:          req.Price,
		Organizer:      req.Organizer,
		ImageURL:       req.ImageURL,
		AvailableSpots: req.AvailableSpots,
		Contact:        req.Contact,
		CreatedBy:      userID,
	}

	if err := s.repo.Create(event); err != nil {
		return nil, err
	}
	return event, nil
}

func (s *Service) Update(id int, req *UpdateEventRequest) error {
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
	if req.Date != "" {
		updates["date"] = req.Date
	}
	if req.Time != "" {
		updates["time"] = req.Time
	}
	if req.Type != "" {
		updates["type"] = req.Type
	}
	if req.Category != "" {
		updates["category"] = req.Category
	}
	if req.Location != "" {
		updates["location"] = req.Location
	}
	if req.Price != "" {
		updates["price"] = req.Price
	}
	if req.Organizer != "" {
		updates["organizer"] = req.Organizer
	}
	if req.ImageURL != "" {
		updates["image_url"] = req.ImageURL
	}
	if req.AvailableSpots > 0 {
		updates["available_spots"] = req.AvailableSpots
	}
	if req.Contact != "" {
		updates["contact"] = req.Contact
	}

	return s.repo.Update(id, updates)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *Service) GetByID(id int, userID int) (*EventResponse, error) {
	event, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	registrationsCount, _ := s.repo.GetRegistrationsCount(id)
	isRegistered, _ := s.repo.IsRegistered(id, userID)

	return &EventResponse{
		ID:                 event.ID,
		Title:              event.Title,
		ShortText:          event.ShortText,
		FullText:           event.FullText,
		Date:               event.Date,
		Time:               event.Time,
		Type:               event.Type,
		Category:           event.Category,
		Location:           event.Location,
		Price:              event.Price,
		Organizer:          event.Organizer,
		ImageURL:           event.ImageURL,
		AvailableSpots:     event.AvailableSpots,
		Contact:            event.Contact,
		RegistrationsCount: registrationsCount,
		IsRegistered:       isRegistered,
	}, nil
}

func (s *Service) GetAll(query *GetEventsQuery, userID int) ([]EventResponse, error) {
	limit := query.Limit
	if limit == 0 {
		limit = 20
	}

	events, err := s.repo.GetAll(query.Type, query.Category, query.DateFrom, query.DateTo, limit, query.Offset)
	if err != nil {
		return nil, err
	}

	var response []EventResponse
	for _, event := range events {
		registrationsCount, _ := s.repo.GetRegistrationsCount(event.ID)
		isRegistered, _ := s.repo.IsRegistered(event.ID, userID)

		response = append(response, EventResponse{
			ID:                 event.ID,
			Title:              event.Title,
			ShortText:          event.ShortText,
			Date:               event.Date,
			Time:               event.Time,
			Type:               event.Type,
			Category:           event.Category,
			Location:           event.Location,
			Price:              event.Price,
			Organizer:          event.Organizer,
			ImageURL:           event.ImageURL,
			AvailableSpots:     event.AvailableSpots,
			RegistrationsCount: registrationsCount,
			IsRegistered:       isRegistered,
		})
	}
	return response, nil
}

func (s *Service) Register(eventID, userID int) error {
	event, err := s.repo.FindByID(eventID)
	if err != nil {
		return err
	}

	// Проверяем количество мест
	registrationsCount, _ := s.repo.GetRegistrationsCount(eventID)
	if event.AvailableSpots > 0 && registrationsCount >= event.AvailableSpots {
		return errors.New("нет свободных мест")
	}

	// Проверяем, не зарегистрирован ли уже
	isRegistered, _ := s.repo.IsRegistered(eventID, userID)
	if isRegistered {
		return errors.New("вы уже зарегистрированы на это мероприятие")
	}

	return s.repo.AddRegistration(eventID, userID)
}

func (s *Service) Unregister(eventID, userID int) error {
	return s.repo.RemoveRegistration(eventID, userID)
}
