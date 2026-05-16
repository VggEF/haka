package skilltree

import (
	"database/sql"
	"fmt"
	"strings"
	"student-app/internal/models"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// ========== Навыки ==========
func (r *Repository) CreateSkill(skill *models.Skill) error {
	query := `
        INSERT INTO skills (name, description, xp_cost, icon, category, parent_skill_id, required_skill_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id
    `
	return r.db.QueryRow(query, skill.Name, skill.Description, skill.XPCost,
		skill.Icon, skill.Category, skill.ParentSkillID, skill.RequiredSkillID).Scan(&skill.ID)
}

func (r *Repository) UpdateSkill(id int, updates map[string]interface{}) error {
	setParts := make([]string, 0, len(updates))
	args := make([]interface{}, 0, len(updates)+1)
	i := 1

	for key, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE skills SET %s WHERE id = $%d", strings.Join(setParts, ", "), i)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Repository) DeleteSkill(id int) error {
	query := `DELETE FROM skills WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) FindSkillByID(id int) (*models.Skill, error) {
	var skill models.Skill
	query := `
        SELECT id, name, description, xp_cost, icon, category, parent_skill_id, required_skill_id
        FROM skills
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(&skill.ID, &skill.Name, &skill.Description,
		&skill.XPCost, &skill.Icon, &skill.Category, &skill.ParentSkillID, &skill.RequiredSkillID)
	if err != nil {
		return nil, err
	}
	return &skill, nil
}

func (r *Repository) GetAllSkills(category string) ([]models.Skill, error) {
	query := `
        SELECT id, name, description, xp_cost, icon, category, parent_skill_id, required_skill_id
        FROM skills
        WHERE 1=1
    `
	args := []interface{}{}
	i := 1

	if category != "" {
		query += fmt.Sprintf(" AND category = $%d", i)
		args = append(args, category)
		i++
	}

	query += " ORDER BY xp_cost ASC"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skills []models.Skill
	for rows.Next() {
		var skill models.Skill
		err := rows.Scan(&skill.ID, &skill.Name, &skill.Description, &skill.XPCost,
			&skill.Icon, &skill.Category, &skill.ParentSkillID, &skill.RequiredSkillID)
		if err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}
	return skills, nil
}

// ========== Пользовательские навыки ==========
func (r *Repository) UnlockSkill(userID, skillID int) error {
	query := `
        INSERT INTO user_skills (user_id, skill_id)
        VALUES ($1, $2)
        ON CONFLICT (user_id, skill_id) DO NOTHING
    `
	_, err := r.db.Exec(query, userID, skillID)
	return err
}

func (r *Repository) IsSkillUnlocked(userID, skillID int) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM user_skills WHERE user_id = $1 AND skill_id = $2)`
	err := r.db.QueryRow(query, userID, skillID).Scan(&exists)
	return exists, err
}

func (r *Repository) GetUserSkills(userID int, category string) ([]models.Skill, error) {
	query := `
        SELECT s.id, s.name, s.description, s.xp_cost, s.icon, s.category, s.parent_skill_id, s.required_skill_id
        FROM skills s
        JOIN user_skills us ON s.id = us.skill_id
        WHERE us.user_id = $1
    `
	args := []interface{}{userID}
	i := 2

	if category != "" {
		query += fmt.Sprintf(" AND s.category = $%d", i)
		args = append(args, category)
		i++
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skills []models.Skill
	for rows.Next() {
		var skill models.Skill
		err := rows.Scan(&skill.ID, &skill.Name, &skill.Description, &skill.XPCost,
			&skill.Icon, &skill.Category, &skill.ParentSkillID, &skill.RequiredSkillID)
		if err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}
	return skills, nil
}

func (r *Repository) GetUserSkillUnlockTime(userID, skillID int) (*time.Time, error) {
	var unlockedAt time.Time
	query := `SELECT unlocked_at FROM user_skills WHERE user_id = $1 AND skill_id = $2`
	err := r.db.QueryRow(query, userID, skillID).Scan(&unlockedAt)
	if err != nil {
		return nil, err
	}
	return &unlockedAt, nil
}
