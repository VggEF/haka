package planner

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

// ========== Привычки ==========
func (r *Repository) CreateHabit(habit *models.Habit) error {
	query := `
        INSERT INTO habits (name, description, xp_reward, created_by)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `
	return r.db.QueryRow(query, habit.Name, habit.Description, habit.XPReward, habit.CreatedBy).Scan(&habit.ID, &habit.CreatedAt)
}

func (r *Repository) UpdateHabit(id int, updates map[string]interface{}) error {
	setParts := make([]string, 0, len(updates))
	args := make([]interface{}, 0, len(updates)+1)
	i := 1

	for key, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE habits SET %s WHERE id = $%d", strings.Join(setParts, ", "), i)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Repository) DeleteHabit(id int) error {
	query := `DELETE FROM habits WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) FindHabitByID(id int) (*models.Habit, error) {
	var habit models.Habit
	query := `
        SELECT id, name, description, xp_reward, created_by, created_at
        FROM habits
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(&habit.ID, &habit.Name, &habit.Description,
		&habit.XPReward, &habit.CreatedBy, &habit.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &habit, nil
}

func (r *Repository) GetAllHabits(isActive bool) ([]models.Habit, error) {
	query := `
        SELECT id, name, description, xp_reward, created_at
        FROM habits
        WHERE 1=1
    `
	if isActive {
		query += " AND is_active = true"
	}
	query += " ORDER BY created_at ASC"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var habits []models.Habit
	for rows.Next() {
		var habit models.Habit
		err := rows.Scan(&habit.ID, &habit.Name, &habit.Description,
			&habit.XPReward, &habit.CreatedAt)
		if err != nil {
			return nil, err
		}
		habits = append(habits, habit)
	}
	return habits, nil
}

// ========== Пользовательские привычки ==========
func (r *Repository) GetUserHabits(userID int) (map[int]models.UserHabit, error) {
	query := `
        SELECT habit_id, streak, last_completed
        FROM user_habits
        WHERE user_id = $1
    `
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	habits := make(map[int]models.UserHabit)
	for rows.Next() {
		var uh models.UserHabit
		err := rows.Scan(&uh.HabitID, &uh.Streak, &uh.LastCompleted)
		if err != nil {
			return nil, err
		}
		habits[uh.HabitID] = uh
	}
	return habits, nil
}

func (r *Repository) CompleteHabit(userID, habitID, xpReward int) error {
	// Начинаем транзакцию
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	today := time.Now().Truncate(24 * time.Hour)

	// Проверяем, не выполнялась ли уже сегодня
	var completed bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM habit_logs WHERE user_id = $1 AND habit_id = $2 AND completed_date = $3)`
	err = tx.QueryRow(checkQuery, userID, habitID, today).Scan(&completed)
	if err != nil {
		return err
	}
	if completed {
		return fmt.Errorf("сегодня вы уже выполняли эту привычку")
	}

	// Обновляем streak
	var currentStreak int
	var lastCompleted sql.NullTime
	getUserHabitQuery := `SELECT streak, last_completed FROM user_habits WHERE user_id = $1 AND habit_id = $2`
	err = tx.QueryRow(getUserHabitQuery, userID, habitID).Scan(&currentStreak, &lastCompleted)
	if err == sql.ErrNoRows {
		currentStreak = 0
	} else if err != nil {
		return err
	}

	// Проверяем, не прервана ли серия
	if lastCompleted.Valid {
		lastDate := lastCompleted.Time.Truncate(24 * time.Hour)
		daysDiff := int(today.Sub(lastDate).Hours() / 24)
		if daysDiff == 1 {
			currentStreak++
		} else if daysDiff > 1 {
			currentStreak = 1
		} else {
		}
	} else {
		currentStreak = 1
	}

	// Обновляем user_habits
	upsertQuery := `
        INSERT INTO user_habits (user_id, habit_id, streak, last_completed)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (user_id, habit_id) DO UPDATE SET
            streak = EXCLUDED.streak,
            last_completed = EXCLUDED.last_completed
    `
	_, err = tx.Exec(upsertQuery, userID, habitID, currentStreak, today)
	if err != nil {
		return err
	}

	// Добавляем лог
	logQuery := `
        INSERT INTO habit_logs (user_id, habit_id, completed_date, xp_earned)
        VALUES ($1, $2, $3, $4)
    `
	_, err = tx.Exec(logQuery, userID, habitID, today, xpReward)
	if err != nil {
		return err
	}

	// Начисляем XP
	xpQuery := `UPDATE users SET total_xp = total_xp + $1 WHERE id = $2`
	_, err = tx.Exec(xpQuery, xpReward, userID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *Repository) GetHabitLogs(userID, habitID int, limit int) ([]models.HabitLog, error) {
	query := `
        SELECT id, user_id, habit_id, completed_date, xp_earned
        FROM habit_logs
        WHERE user_id = $1
    `
	args := []interface{}{userID}
	i := 2

	if habitID > 0 {
		query += fmt.Sprintf(" AND habit_id = $%d", i)
		args = append(args, habitID)
		i++
	}

	if limit == 0 {
		limit = 30
	}
	query += fmt.Sprintf(" ORDER BY completed_date DESC LIMIT $%d", i)
	args = append(args, limit)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []models.HabitLog
	for rows.Next() {
		var log models.HabitLog
		err := rows.Scan(&log.ID, &log.UserID, &log.HabitID, &log.CompletedDate, &log.XPEarned)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

func (r *Repository) GetWeeklyStats(userID int) (map[string]int, error) {
	stats := make(map[string]int)

	query := `
        SELECT EXTRACT(DOW FROM completed_date) as day_of_week, COUNT(*)
        FROM habit_logs
        WHERE user_id = $1 AND completed_date >= CURRENT_DATE - INTERVAL '7 days'
        GROUP BY EXTRACT(DOW FROM completed_date)
    `
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var day int
		var count int
		err := rows.Scan(&day, &count)
		if err != nil {
			return nil, err
		}
		dayNames := map[int]string{
			0: "Вс", 1: "Пн", 2: "Вт", 3: "Ср", 4: "Чт", 5: "Пт", 6: "Сб",
		}
		stats[dayNames[day]] = count
	}
	return stats, nil
}

func (r *Repository) AddXP(userID, xp int) error {
	query := `UPDATE users SET total_xp = total_xp + $1 WHERE id = $2`
	_, err := r.db.Exec(query, xp, userID)
	return err
}
