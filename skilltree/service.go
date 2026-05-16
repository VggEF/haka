package skilltree

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

func (s *Service) CreateSkill(req *CreateSkillRequest) (*models.Skill, error) {
	skill := &models.Skill{
		Name:            req.Name,
		Description:     req.Description,
		XPCost:          req.XPCost,
		Icon:            req.Icon,
		Category:        req.Category,
		ParentSkillID:   req.ParentSkillID,
		RequiredSkillID: req.RequiredSkillID,
	}

	if skill.Icon == "" {
		skill.Icon = "📚"
	}

	if err := s.repo.CreateSkill(skill); err != nil {
		return nil, err
	}
	return skill, nil
}

func (s *Service) UpdateSkill(id int, req *UpdateSkillRequest) error {
	updates := make(map[string]interface{})

	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.XPCost > 0 {
		updates["xp_cost"] = req.XPCost
	}
	if req.Icon != "" {
		updates["icon"] = req.Icon
	}
	if req.Category != "" {
		updates["category"] = req.Category
	}
	if req.ParentSkillID != nil {
		updates["parent_skill_id"] = req.ParentSkillID
	}
	if req.RequiredSkillID != nil {
		updates["required_skill_id"] = req.RequiredSkillID
	}

	return s.repo.UpdateSkill(id, updates)
}

func (s *Service) DeleteSkill(id int) error {
	return s.repo.DeleteSkill(id)
}

func (s *Service) GetAllSkills(userID int, category string) ([]SkillResponse, error) {
	skills, err := s.repo.GetAllSkills(category)
	if err != nil {
		return nil, err
	}

	// Получаем XP пользователя
	userXP, err := s.getUserXP(userID)
	if err != nil {
		userXP = 0
	}

	var response []SkillResponse
	for _, skill := range skills {
		isUnlocked, _ := s.repo.IsSkillUnlocked(userID, skill.ID)

		// Проверяем, можно ли открыть навык
		canUnlock := false
		if !isUnlocked {
			canUnlock = true

			// Проверяем достаточно ли XP
			if userXP < skill.XPCost {
				canUnlock = false
			}

			// Проверяем требования по родительским навыкам
			if skill.RequiredSkillID != nil {
				requiredUnlocked, _ := s.repo.IsSkillUnlocked(userID, *skill.RequiredSkillID)
				if !requiredUnlocked {
					canUnlock = false
				}
			}
		}

		var unlockedAt *time.Time
		if isUnlocked {
			ut, _ := s.repo.GetUserSkillUnlockTime(userID, skill.ID)
			unlockedAt = ut
		}

		response = append(response, SkillResponse{
			ID:              skill.ID,
			Name:            skill.Name,
			Description:     skill.Description,
			XPCost:          skill.XPCost,
			Icon:            skill.Icon,
			Category:        skill.Category,
			ParentSkillID:   skill.ParentSkillID,
			RequiredSkillID: skill.RequiredSkillID,
			IsUnlocked:      isUnlocked,
			CanUnlock:       canUnlock,
			UnlockedAt:      unlockedAt,
		})
	}
	return response, nil
}

func (s *Service) UnlockSkill(userID, skillID int) error {
	// Проверяем, не открыт ли уже
	isUnlocked, err := s.repo.IsSkillUnlocked(userID, skillID)
	if err != nil {
		return err
	}
	if isUnlocked {
		return errors.New("навык уже открыт")
	}

	// Получаем информацию о навыке
	skill, err := s.repo.FindSkillByID(skillID)
	if err != nil {
		return err
	}

	// Проверяем достаточно ли XP
	userXP, err := s.getUserXP(userID)
	if err != nil {
		return errors.New("не удалось получить XP пользователя")
	}
	if userXP < skill.XPCost {
		return errors.New("недостаточно XP для открытия навыка")
	}

	// Проверяем требования по родительским навыкам
	if skill.RequiredSkillID != nil {
		requiredUnlocked, _ := s.repo.IsSkillUnlocked(userID, *skill.RequiredSkillID)
		if !requiredUnlocked {
			return errors.New("требуется открыть prerequisite навык")
		}
	}

	// Открываем навык
	if err := s.repo.UnlockSkill(userID, skillID); err != nil {
		return err
	}

	// Списываем XP
	if err := s.spendXP(userID, skill.XPCost); err != nil {
		return err
	}

	return nil
}

func (s *Service) getUserXP(userID int) (int, error) {
	var xp int
	query := `SELECT total_xp FROM users WHERE id = $1`
	err := s.repo.db.QueryRow(query, userID).Scan(&xp)
	return xp, err
}

func (s *Service) spendXP(userID, amount int) error {
	query := `UPDATE users SET total_xp = total_xp - $1 WHERE id = $2 AND total_xp >= $1`
	result, err := s.repo.db.Exec(query, amount, userID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("недостаточно XP")
	}
	return nil
}
