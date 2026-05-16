package achievements

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"student-app/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// ========== Ачивки ==========
func (r *Repository) CreateAchievement(ach *models.Achievement) error {
	query := `
        INSERT INTO achievements (title, description, xp, icon, rarity, category, condition_text, created_by)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, created_at
    `
	return r.db.QueryRow(query, ach.Title, ach.Description, ach.XP, ach.Icon,
		ach.Rarity, ach.Category, ach.Condition, ach.CreatedBy).Scan(&ach.ID, &ach.CreatedAt)
}

func (r *Repository) UpdateAchievement(id int, updates map[string]interface{}) error {
	setParts := make([]string, 0, len(updates))
	args := make([]interface{}, 0, len(updates)+1)
	i := 1

	for key, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE achievements SET %s WHERE id = $%d", strings.Join(setParts, ", "), i)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Repository) DeleteAchievement(id int) error {
	query := `DELETE FROM achievements WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) FindAchievementByID(id int) (*models.Achievement, error) {
	var ach models.Achievement
	query := `
        SELECT id, title, description, xp, icon, rarity, category, condition_text, created_by, created_at
        FROM achievements
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(&ach.ID, &ach.Title, &ach.Description,
		&ach.XP, &ach.Icon, &ach.Rarity, &ach.Category, &ach.Condition,
		&ach.CreatedBy, &ach.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &ach, nil
}

func (r *Repository) GetAllAchievements(category string, limit, offset int) ([]models.Achievement, error) {
	log.Printf("🔍 SQL запрос: category=%s, limit=%d, offset=%d", category, limit, offset)

	query := `
        SELECT id, title, description, xp, icon, rarity, category, condition_text, created_at
        FROM achievements
        WHERE 1=1
    `
	args := []interface{}{}
	i := 1

	if category != "" {
		query += fmt.Sprintf(" AND category = $%d", i)
		args = append(args, category)
		i++
	}

	query += fmt.Sprintf(" ORDER BY xp DESC LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)

	log.Printf("📝 SQL: %s", query)
	log.Printf("📝 Args: %v", args)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		log.Printf("❌ Ошибка запроса: %v", err)
		return nil, err
	}
	defer rows.Close()

	var achievements []models.Achievement
	for rows.Next() {
		var ach models.Achievement
		err := rows.Scan(&ach.ID, &ach.Title, &ach.Description, &ach.XP,
			&ach.Icon, &ach.Rarity, &ach.Category, &ach.Condition, &ach.CreatedAt)
		if err != nil {
			log.Printf("❌ Ошибка сканирования: %v", err)
			return nil, err
		}
		achievements = append(achievements, ach)
		log.Printf("📦 Найдена ачивка: ID=%d, Title=%s", ach.ID, ach.Title)
	}

	log.Printf("✅ Всего найдено: %d", len(achievements))
	return achievements, nil
}

// ========== Пользовательские ачивки ==========
func (r *Repository) UnlockAchievement(userID, achievementID int) error {
	query := `
        INSERT INTO user_achievements (user_id, achievement_id)
        VALUES ($1, $2)
        ON CONFLICT (user_id, achievement_id) DO NOTHING
    `
	_, err := r.db.Exec(query, userID, achievementID)
	return err
}

func (r *Repository) IsAchievementUnlocked(userID, achievementID int) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM user_achievements WHERE user_id = $1 AND achievement_id = $2)`
	err := r.db.QueryRow(query, userID, achievementID).Scan(&exists)
	return exists, err
}

func (r *Repository) GetUserAchievements(userID int, category string, unlocked bool, limit, offset int) ([]models.Achievement, error) {
	query := `
        SELECT a.id, a.title, a.description, a.xp, a.icon, a.rarity, a.category, a.condition_text, ua.earned_at
        FROM achievements a
        LEFT JOIN user_achievements ua ON a.id = ua.achievement_id AND ua.user_id = $1
        WHERE 1=1
    `
	args := []interface{}{userID}
	i := 2

	if category != "" {
		query += fmt.Sprintf(" AND a.category = $%d", i)
		args = append(args, category)
		i++
	}
	if unlocked {
		query += " AND ua.user_id IS NOT NULL"
	} else {
		query += " AND ua.user_id IS NULL"
	}

	query += fmt.Sprintf(" ORDER BY a.xp DESC LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var achievements []models.Achievement
	for rows.Next() {
		var ach models.Achievement
		var earnedAt sql.NullTime
		err := rows.Scan(&ach.ID, &ach.Title, &ach.Description, &ach.XP,
			&ach.Icon, &ach.Rarity, &ach.Category, &ach.Condition, &earnedAt)
		if err != nil {
			return nil, err
		}
		achievements = append(achievements, ach)
	}
	return achievements, nil
}

func (r *Repository) GetUserTotalXP(userID int) (int, error) {
	query := `
        SELECT COALESCE(SUM(a.xp), 0)
        FROM user_achievements ua
        JOIN achievements a ON ua.achievement_id = a.id
        WHERE ua.user_id = $1
    `
	var totalXP int
	err := r.db.QueryRow(query, userID).Scan(&totalXP)
	return totalXP, err
}

func (r *Repository) UpdateUserTotalXP(userID, xp int) error {
	query := `UPDATE users SET total_xp = total_xp + $1 WHERE id = $2`
	_, err := r.db.Exec(query, xp, userID)
	return err
}

// ========== Ежедневные квесты ==========
func (r *Repository) CreateDailyQuest(quest *models.DailyQuest) error {
	query := `
        INSERT INTO daily_quests (title, description, xp, coins, requirement, expires_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `
	return r.db.QueryRow(query, quest.Title, quest.Description, quest.XP,
		quest.Coins, quest.Requirement, quest.ExpiresAt).Scan(&quest.ID)
}

func (r *Repository) GetAllDailyQuests() ([]models.DailyQuest, error) {
	query := `
        SELECT id, title, description, xp, coins, requirement, expires_at
        FROM daily_quests
        WHERE expires_at > NOW()
        ORDER BY expires_at ASC
    `
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quests []models.DailyQuest
	for rows.Next() {
		var quest models.DailyQuest
		err := rows.Scan(&quest.ID, &quest.Title, &quest.Description,
			&quest.XP, &quest.Coins, &quest.Requirement, &quest.ExpiresAt)
		if err != nil {
			return nil, err
		}
		quests = append(quests, quest)
	}
	return quests, nil
}

func (r *Repository) CompleteQuest(userID, questID int) error {
	query := `
        INSERT INTO user_quests (user_id, quest_id)
        VALUES ($1, $2)
        ON CONFLICT (user_id, quest_id) DO NOTHING
    `
	_, err := r.db.Exec(query, userID, questID)
	return err
}

func (r *Repository) IsQuestCompleted(userID, questID int) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM user_quests WHERE user_id = $1 AND quest_id = $2)`
	err := r.db.QueryRow(query, userID, questID).Scan(&exists)
	return exists, err
}
