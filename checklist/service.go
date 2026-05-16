package checklist

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

// ========== Управление элементами чек-листа (только админ) ==========
func (s *Service) CreateItem(req *CreateChecklistItemRequest) (*models.ChecklistItem, error) {
	item := &models.ChecklistItem{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		XP:          req.XP,
		SortOrder:   req.SortOrder,
		IsActive:    true,
	}

	if item.Category == "" {
		item.Category = "Основные"
	}
	if item.XP == 0 {
		item.XP = 10
	}

	if err := s.repo.CreateItem(item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) UpdateItem(id int, req *UpdateChecklistItemRequest) error {
	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Category != "" {
		updates["category"] = req.Category
	}
	if req.XP > 0 {
		updates["xp"] = req.XP
	}
	if req.SortOrder > 0 {
		updates["sort_order"] = req.SortOrder
	}
	updates["is_active"] = req.IsActive

	return s.repo.UpdateItem(id, updates)
}

func (s *Service) DeleteItem(id int) error {
	return s.repo.DeleteItem(id)
}

// ========== Получение чек-листа для пользователя ==========
func (s *Service) GetChecklist(userID int, query *GetChecklistQuery) ([]ChecklistItemResponse, error) {
	// Получаем все активные элементы
	items, err := s.repo.GetAllItems(query.Category, true)
	if err != nil {
		return nil, err
	}

	// Получаем прогресс пользователя
	progress, err := s.repo.GetUserProgress(userID)
	if err != nil {
		return nil, err
	}

	var response []ChecklistItemResponse
	for _, item := range items {
		isCompleted := progress[item.ID]

		// Фильтр по статусу выполнения
		if query.Completed && !isCompleted {
			continue
		}
		if !query.Completed && query.Completed != false {
			// если нужно показать только невыполненные
		}

		var completedAt *time.Time
		if isCompleted {
			// можно получить время выполнения отдельным запросом
		}

		response = append(response, ChecklistItemResponse{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Category:    item.Category,
			XP:          item.XP,
			SortOrder:   item.SortOrder,
			IsActive:    item.IsActive,
			IsCompleted: isCompleted,
			CompletedAt: completedAt,
		})
	}
	return response, nil
}

func (s *Service) CompleteItem(userID, itemID int) error {
	// Проверяем, существует ли элемент
	item, err := s.repo.FindItemByID(itemID)
	if err != nil {
		return errors.New("элемент не найден")
	}
	if !item.IsActive {
		return errors.New("элемент неактивен")
	}

	// Проверяем, не выполнен ли уже
	isCompleted, err := s.repo.IsItemCompleted(userID, itemID)
	if err != nil {
		return err
	}
	if isCompleted {
		return errors.New("элемент уже выполнен")
	}

	// Отмечаем выполнение
	if err := s.repo.CompleteItem(userID, itemID); err != nil {
		return err
	}

	// Начисляем XP пользователю
	if err := s.repo.UpdateUserTotalXP(userID, item.XP); err != nil {
		// Логируем ошибку, но не откатываем выполнение
	}

	return nil
}

func (s *Service) GetUserProgress(userID int) (completedCount int, totalCount int, totalXP int, err error) {
	// Получаем все активные элементы
	items, err := s.repo.GetAllItems("", true)
	if err != nil {
		return 0, 0, 0, err
	}
	totalCount = len(items)

	// Получаем выполненные
	completed, err := s.repo.GetUserCompletedItems(userID)
	if err != nil {
		return 0, 0, 0, err
	}
	completedCount = len(completed)

	// Получаем общее XP
	totalXP, err = s.repo.GetUserTotalCompletedXP(userID)
	if err != nil {
		return 0, 0, 0, err
	}

	return completedCount, totalCount, totalXP, nil
}

func (s *Service) GetCategories() ([]string, error) {
	items, err := s.repo.GetAllItems("", true)
	if err != nil {
		return nil, err
	}

	categoriesMap := make(map[string]bool)
	for _, item := range items {
		if item.Category != "" {
			categoriesMap[item.Category] = true
		}
	}

	categories := make([]string, 0, len(categoriesMap))
	for cat := range categoriesMap {
		categories = append(categories, cat)
	}
	return categories, nil
}
