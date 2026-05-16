package achievements

import (
	"errors"
	"student-app/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// ========== Ачивки ==========
func (s *Service) CreateAchievement(req *CreateAchievementRequest, userID int) (*models.Achievement, error) {
	ach := &models.Achievement{
		Title:       req.Title,
		Description: req.Description,
		XP:          req.XP,
		Icon:        req.Icon,
		Rarity:      req.Rarity,
		Category:    req.Category,
		Condition:   req.Condition,
		CreatedBy:   userID,
	}

	if ach.Rarity == "" {
		ach.Rarity = "common"
	}
	if ach.Icon == "" {
		ach.Icon = "🏆"
	}

	if err := s.repo.CreateAchievement(ach); err != nil {
		return nil, err
	}
	return ach, nil
}

func (s *Service) UpdateAchievement(id int, req *UpdateAchievementRequest) error {
	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.XP > 0 {
		updates["xp"] = req.XP
	}
	if req.Icon != "" {
		updates["icon"] = req.Icon
	}
	if req.Rarity != "" {
		updates["rarity"] = req.Rarity
	}
	if req.Category != "" {
		updates["category"] = req.Category
	}
	if req.Condition != "" {
		updates["condition_text"] = req.Condition
	}

	return s.repo.UpdateAchievement(id, updates)
}

func (s *Service) DeleteAchievement(id int) error {
	return s.repo.DeleteAchievement(id)
}

func (s *Service) GetAllAchievements(category string, limit, offset int) ([]AchievementResponse, error) {
	achievements, err := s.repo.GetAllAchievements(category, limit, offset)
	if err != nil {
		return nil, err
	}

	var response []AchievementResponse
	for _, ach := range achievements {
		response = append(response, AchievementResponse{
			ID:          ach.ID,
			Title:       ach.Title,
			Description: ach.Description,
			XP:          ach.XP,
			Icon:        ach.Icon,
			Rarity:      ach.Rarity,
			Category:    ach.Category,
			Condition:   ach.Condition,
			CreatedAt:   ach.CreatedAt,
			IsUnlocked:  false,
		})
	}
	return response, nil
}

// ========== Выдача ачивок ==========
func (s *Service) AwardAchievement(userID, achievementID, awardedBy int) error {
	// Проверяем, не получена ли уже
	isUnlocked, err := s.repo.IsAchievementUnlocked(userID, achievementID)
	if err != nil {
		return err
	}
	if isUnlocked {
		return errors.New("ачивка уже получена пользователем")
	}

	// Получаем ачивку для XP
	achievement, err := s.repo.FindAchievementByID(achievementID)
	if err != nil {
		return err
	}

	// Выдаем ачивку
	if err := s.repo.UnlockAchievement(userID, achievementID); err != nil {
		return err
	}

	// Начисляем XP пользователю
	if err := s.repo.UpdateUserTotalXP(userID, achievement.XP); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUserAchievements(userID int, query *GetUserAchievementsQuery) ([]AchievementResponse, error) {
	limit := query.Limit
	if limit == 0 {
		limit = 50
	}

	achievements, err := s.repo.GetUserAchievements(userID, query.Category, query.Unlocked, limit, query.Offset)
	if err != nil {
		return nil, err
	}

	var response []AchievementResponse
	for _, ach := range achievements {
		isUnlocked, _ := s.repo.IsAchievementUnlocked(userID, ach.ID)
		response = append(response, AchievementResponse{
			ID:          ach.ID,
			Title:       ach.Title,
			Description: ach.Description,
			XP:          ach.XP,
			Icon:        ach.Icon,
			Rarity:      ach.Rarity,
			Category:    ach.Category,
			Condition:   ach.Condition,
			CreatedAt:   ach.CreatedAt,
			IsUnlocked:  isUnlocked,
		})
	}
	return response, nil
}

func (s *Service) GetUserTotalXP(userID int) (int, error) {
	return s.repo.GetUserTotalXP(userID)
}

// ========== Ежедневные квесты ==========
func (s *Service) CreateDailyQuest(req *CreateDailyQuestRequest) (*models.DailyQuest, error) {
	quest := &models.DailyQuest{
		Title:       req.Title,
		Description: req.Description,
		XP:          req.XP,
		Coins:       req.Coins,
		Requirement: req.Requirement,
		ExpiresAt:   req.ExpiresAt,
	}

	if err := s.repo.CreateDailyQuest(quest); err != nil {
		return nil, err
	}
	return quest, nil
}

func (s *Service) GetDailyQuests(userID int) ([]DailyQuestResponse, error) {
	quests, err := s.repo.GetAllDailyQuests()
	if err != nil {
		return nil, err
	}

	var response []DailyQuestResponse
	for _, quest := range quests {
		isCompleted, _ := s.repo.IsQuestCompleted(userID, quest.ID)
		response = append(response, DailyQuestResponse{
			ID:          quest.ID,
			Title:       quest.Title,
			Description: quest.Description,
			XP:          quest.XP,
			Coins:       quest.Coins,
			Requirement: quest.Requirement,
			ExpiresAt:   quest.ExpiresAt,
			IsCompleted: isCompleted,
		})
	}
	return response, nil
}

func (s *Service) CompleteQuest(userID, questID int) error {
	// Проверяем, не выполнен ли уже
	isCompleted, err := s.repo.IsQuestCompleted(userID, questID)
	if err != nil {
		return err
	}
	if isCompleted {
		return errors.New("квест уже выполнен")
	}

	// Получаем квест для награды
	quests, err := s.repo.GetAllDailyQuests()
	if err != nil {
		return err
	}

	var targetQuest *models.DailyQuest
	for _, q := range quests {
		if q.ID == questID {
			targetQuest = &q
			break
		}
	}
	if targetQuest == nil {
		return errors.New("квест не найден")
	}

	// Отмечаем выполнение
	if err := s.repo.CompleteQuest(userID, questID); err != nil {
		return err
	}

	// Начисляем награду
	if err := s.repo.UpdateUserTotalXP(userID, targetQuest.XP); err != nil {
		return err
	}

	return nil
}
