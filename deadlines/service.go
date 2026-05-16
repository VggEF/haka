package deadlines

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

func (s *Service) Create(req *CreateDeadlineRequest, userID int) (*models.Deadline, error) {
	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		return nil, errors.New("неверный формат даты")
	}

	priority := req.Priority
	if priority == "" {
		priority = "medium"
	}

	deadline := &models.Deadline{
		Title:       req.Title,
		Subject:     req.Subject,
		GroupName:   req.GroupName,
		DueDate:     dueDate,
		DueTime:     req.DueTime,
		Priority:    priority,
		Description: req.Description,
		Status:      "pending",
		CreatedBy:   userID,
		AssignedTo:  req.AssignedTo,
	}

	if err := s.repo.Create(deadline); err != nil {
		return nil, err
	}
	return deadline, nil
}

func (s *Service) Update(id int, req *UpdateDeadlineRequest) error {
	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Subject != "" {
		updates["subject"] = req.Subject
	}
	if req.GroupName != "" {
		updates["group_name"] = req.GroupName
	}
	if req.DueDate != "" {
		updates["due_date"] = req.DueDate
	}
	if req.DueTime != "" {
		updates["due_time"] = req.DueTime
	}
	if req.Priority != "" {
		updates["priority"] = req.Priority
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}

	return s.repo.Update(id, updates)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *Service) GetByID(id int) (*DeadlineResponse, error) {
	deadline, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	daysLeft := int(time.Until(deadline.DueDate).Hours() / 24)
	isOverdue := deadline.Status == "pending" && deadline.DueDate.Before(time.Now())

	return &DeadlineResponse{
		ID:          deadline.ID,
		Title:       deadline.Title,
		Subject:     deadline.Subject,
		GroupName:   deadline.GroupName,
		DueDate:     deadline.DueDate,
		DueTime:     deadline.DueTime,
		Priority:    deadline.Priority,
		Description: deadline.Description,
		Status:      deadline.Status,
		CreatedBy:   deadline.CreatedBy,
		CreatedAt:   deadline.CreatedAt,
		DaysLeft:    daysLeft,
		IsOverdue:   isOverdue,
	}, nil
}

func (s *Service) GetUserDeadlines(userID int, query *GetDeadlinesQuery) ([]DeadlineResponse, error) {
	deadlines, err := s.repo.FindByUser(userID, query)
	if err != nil {
		return nil, err
	}

	var response []DeadlineResponse
	for _, d := range deadlines {
		daysLeft := int(time.Until(d.DueDate).Hours() / 24)
		isOverdue := d.Status == "pending" && d.DueDate.Before(time.Now())

		response = append(response, DeadlineResponse{
			ID:          d.ID,
			Title:       d.Title,
			Subject:     d.Subject,
			GroupName:   d.GroupName,
			DueDate:     d.DueDate,
			DueTime:     d.DueTime,
			Priority:    d.Priority,
			Description: d.Description,
			Status:      d.Status,
			CreatedBy:   d.CreatedBy,
			CreatedAt:   d.CreatedAt,
			DaysLeft:    daysLeft,
			IsOverdue:   isOverdue,
		})
	}
	return response, nil
}

func (s *Service) GetGroupDeadlines(groupName string, query *GetDeadlinesQuery) ([]DeadlineResponse, error) {
	deadlines, err := s.repo.FindByGroup(groupName, query)
	if err != nil {
		return nil, err
	}

	var response []DeadlineResponse
	for _, d := range deadlines {
		daysLeft := int(time.Until(d.DueDate).Hours() / 24)
		isOverdue := d.Status == "pending" && d.DueDate.Before(time.Now())

		response = append(response, DeadlineResponse{
			ID:          d.ID,
			Title:       d.Title,
			Subject:     d.Subject,
			GroupName:   d.GroupName,
			DueDate:     d.DueDate,
			DueTime:     d.DueTime,
			Priority:    d.Priority,
			Description: d.Description,
			Status:      d.Status,
			CreatedBy:   d.CreatedBy,
			CreatedAt:   d.CreatedAt,
			DaysLeft:    daysLeft,
			IsOverdue:   isOverdue,
		})
	}
	return response, nil
}

func (s *Service) CompleteDeadline(id, userID int) error {
	deadline, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Проверяем, принадлежит ли дедлайн пользователю
	if deadline.AssignedTo != userID && deadline.AssignedTo != 0 {
		return errors.New("дедлайн не принадлежит вам")
	}

	return s.repo.UpdateStatus(id, "completed")
}

func (s *Service) GetUpcomingDeadlines(userID int) ([]DeadlineResponse, error) {
	deadlines, err := s.repo.GetUpcomingDeadlines(userID, 10)
	if err != nil {
		return nil, err
	}

	var response []DeadlineResponse
	for _, d := range deadlines {
		daysLeft := int(time.Until(d.DueDate).Hours() / 24)
		response = append(response, DeadlineResponse{
			ID:        d.ID,
			Title:     d.Title,
			Subject:   d.Subject,
			DueDate:   d.DueDate,
			DueTime:   d.DueTime,
			Priority:  d.Priority,
			DaysLeft:  daysLeft,
			IsOverdue: d.DueDate.Before(time.Now()),
		})
	}
	return response, nil
}
