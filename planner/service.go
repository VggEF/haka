package planner

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

// ========== Управление привычками (админ) ==========
func (s *Service) CreateHabit(req *CreateHabitRequest, userID int) (*models.Habit, error) {
	habit := &models.Habit{
		Name:        req.Name,
		Description: req.Description,
		XPReward:    req.XPReward,
		CreatedBy:   userID,
	}

	if habit.XPReward == 0 {
		habit.XPReward = 10
	}

	if err := s.repo.CreateHabit(habit); err != nil {
		return nil, err
	}
	return habit, nil
}

func (s *Service) UpdateHabit(id int, req *UpdateHabitRequest) error {
	updates := make(map[string]interface{})

	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.XPReward > 0 {
		updates["xp_reward"] = req.XPReward
	}
	updates["is_active"] = req.IsActive

	return s.repo.UpdateHabit(id, updates)
}

func (s *Service) DeleteHabit(id int) error {
	return s.repo.DeleteHabit(id)
}

func (s *Service) GetAllHabits(isActive bool) ([]HabitResponse, error) {
	habits, err := s.repo.GetAllHabits(isActive)
	if err != nil {
		return nil, err
	}

	var response []HabitResponse
	for _, h := range habits {
		response = append(response, HabitResponse{
			ID:          h.ID,
			Name:        h.Name,
			Description: h.Description,
			XPReward:    h.XPReward,
			IsActive:    true,
			CreatedAt:   h.CreatedAt,
		})
	}
	return response, nil
}

// ========== Пользовательские операции ==========
func (s *Service) GetUserHabits(userID int) ([]HabitResponse, error) {
	// Получаем все активные привычки
	habits, err := s.repo.GetAllHabits(true)
	if err != nil {
		return nil, err
	}

	// Получаем прогресс пользователя
	userHabits, err := s.repo.GetUserHabits(userID)
	if err != nil {
		return nil, err
	}

	today := time.Now().Truncate(24 * time.Hour)

	var response []HabitResponse
	for _, habit := range habits {
		uh, exists := userHabits[habit.ID]

		isCompleted := false
		if exists && uh.LastCompleted != nil {
			lastDate := uh.LastCompleted.Truncate(24 * time.Hour)
			isCompleted = lastDate.Equal(today)
		}

		response = append(response, HabitResponse{
			ID:            habit.ID,
			Name:          habit.Name,
			Description:   habit.Description,
			XPReward:      habit.XPReward,
			IsActive:      true,
			CreatedAt:     habit.CreatedAt,
			Streak:        uh.Streak,
			IsCompleted:   isCompleted,
			LastCompleted: uh.LastCompleted,
		})
	}
	return response, nil
}

func (s *Service) CompleteHabit(userID, habitID int) error {
	habit, err := s.repo.FindHabitByID(habitID)
	if err != nil {
		return errors.New("привычка не найдена")
	}

	return s.repo.CompleteHabit(userID, habitID, habit.XPReward)
}

func (s *Service) GetHabitLogs(userID, habitID int) ([]HabitLogResponse, error) {
	logs, err := s.repo.GetHabitLogs(userID, habitID, 30)
	if err != nil {
		return nil, err
	}

	var response []HabitLogResponse
	for _, log := range logs {
		// Получаем название привычки
		habit, _ := s.repo.FindHabitByID(log.HabitID)
		habitName := ""
		if habit != nil {
			habitName = habit.Name
		}

		response = append(response, HabitLogResponse{
			ID:            log.ID,
			HabitID:       log.HabitID,
			HabitName:     habitName,
			CompletedDate: log.CompletedDate,
			XPEarned:      log.XPEarned,
		})
	}
	return response, nil
}

func (s *Service) GetWeeklyStats(userID int) (map[string]int, error) {
	return s.repo.GetWeeklyStats(userID)
}
