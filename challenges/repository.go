package challenges

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

// ========== Вызовы ==========
func (r *Repository) CreateChallenge(ch *models.Challenge) error {
	query := `
        INSERT INTO challenges (title, description, type, start_date, end_date, prize_xp, prize_coins, created_by)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, created_at
    `
	return r.db.QueryRow(query, ch.Title, ch.Description, ch.Type,
		ch.StartDate, ch.EndDate, ch.PrizeXP, ch.PrizeCoins, ch.CreatedBy).Scan(&ch.ID, &ch.CreatedAt)
}

func (r *Repository) UpdateChallenge(id int, updates map[string]interface{}) error {
	setParts := make([]string, 0, len(updates))
	args := make([]interface{}, 0, len(updates)+1)
	i := 1

	for key, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE challenges SET %s WHERE id = $%d", strings.Join(setParts, ", "), i)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Repository) DeleteChallenge(id int) error {
	query := `DELETE FROM challenges WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) FindChallengeByID(id int) (*models.Challenge, error) {
	var ch models.Challenge
	query := `
        SELECT id, title, description, type, start_date, end_date, prize_xp, prize_coins, created_by, created_at
        FROM challenges
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(&ch.ID, &ch.Title, &ch.Description,
		&ch.Type, &ch.StartDate, &ch.EndDate, &ch.PrizeXP, &ch.PrizeCoins,
		&ch.CreatedBy, &ch.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &ch, nil
}

func (r *Repository) GetAllChallenges(challengeType, status string, limit, offset int) ([]models.Challenge, error) {
	query := `
        SELECT id, title, description, type, start_date, end_date, prize_xp, prize_coins, created_by, created_at
        FROM challenges
        WHERE 1=1
    `
	args := []interface{}{}
	i := 1
	now := time.Now()

	if challengeType != "" {
		query += fmt.Sprintf(" AND type = $%d", i)
		args = append(args, challengeType)
		i++
	}

	switch status {
	case "active":
		query += fmt.Sprintf(" AND start_date <= $%d AND end_date >= $%d", i, i+1)
		args = append(args, now, now)
		i += 2
	case "upcoming":
		query += fmt.Sprintf(" AND start_date > $%d", i)
		args = append(args, now)
		i++
	case "ended":
		query += fmt.Sprintf(" AND end_date < $%d", i)
		args = append(args, now)
		i++
	}

	if limit == 0 {
		limit = 20
	}
	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var challenges []models.Challenge
	for rows.Next() {
		var ch models.Challenge
		err := rows.Scan(&ch.ID, &ch.Title, &ch.Description, &ch.Type,
			&ch.StartDate, &ch.EndDate, &ch.PrizeXP, &ch.PrizeCoins,
			&ch.CreatedBy, &ch.CreatedAt)
		if err != nil {
			return nil, err
		}
		challenges = append(challenges, ch)
	}
	return challenges, nil
}

// ========== Участники ==========
func (r *Repository) JoinChallenge(challengeID, userID int) error {
	query := `
        INSERT INTO challenge_participants (challenge_id, user_id)
        VALUES ($1, $2)
        ON CONFLICT (challenge_id, user_id) DO NOTHING
    `
	_, err := r.db.Exec(query, challengeID, userID)
	return err
}

func (r *Repository) LeaveChallenge(challengeID, userID int) error {
	query := `DELETE FROM challenge_participants WHERE challenge_id = $1 AND user_id = $2`
	_, err := r.db.Exec(query, challengeID, userID)
	return err
}

func (r *Repository) IsParticipant(challengeID, userID int) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM challenge_participants WHERE challenge_id = $1 AND user_id = $2)`
	err := r.db.QueryRow(query, challengeID, userID).Scan(&exists)
	return exists, err
}

func (r *Repository) UpdateScore(challengeID, userID, score int) error {
	query := `
        UPDATE challenge_participants
        SET score = $3
        WHERE challenge_id = $1 AND user_id = $2
    `
	_, err := r.db.Exec(query, challengeID, userID, score)
	return err
}

func (r *Repository) GetParticipantScore(challengeID, userID int) (int, error) {
	var score int
	query := `SELECT score FROM challenge_participants WHERE challenge_id = $1 AND user_id = $2`
	err := r.db.QueryRow(query, challengeID, userID).Scan(&score)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return score, err
}

func (r *Repository) GetLeaderboard(challengeID int) ([]models.ChallengeParticipant, error) {
	query := `
        SELECT cp.user_id, cp.score, cp.joined_at, u.name
        FROM challenge_participants cp
        JOIN users u ON cp.user_id = u.id
        WHERE cp.challenge_id = $1
        ORDER BY cp.score DESC, cp.joined_at ASC
    `
	rows, err := r.db.Query(query, challengeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []models.ChallengeParticipant
	for rows.Next() {
		var p models.ChallengeParticipant
		var userName string
		err := rows.Scan(&p.UserID, &p.Score, &p.JoinedAt, &userName)
		if err != nil {
			return nil, err
		}
		participants = append(participants, p)
	}
	return participants, nil
}

func (r *Repository) GetParticipantsCount(challengeID int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM challenge_participants WHERE challenge_id = $1`
	err := r.db.QueryRow(query, challengeID).Scan(&count)
	return count, err
}

func (r *Repository) GetParticipantRank(challengeID, userID int) (int, error) {
	query := `
        SELECT COUNT(*) + 1
        FROM challenge_participants
        WHERE challenge_id = $1 AND score > (SELECT score FROM challenge_participants WHERE challenge_id = $1 AND user_id = $2)
    `
	var rank int
	err := r.db.QueryRow(query, challengeID, userID).Scan(&rank)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return rank, err
}

// ========== Награждение ==========
func (r *Repository) AwardPrizes(challengeID int) error {
	// Получаем топ-3 участников
	participants, err := r.GetLeaderboard(challengeID)
	if err != nil {
		return err
	}

	challenge, err := r.FindChallengeByID(challengeID)
	if err != nil {
		return err
	}

	for i, p := range participants {
		if i >= 3 {
			break
		}
		// Начисляем XP и коины
		xpReward := challenge.PrizeXP / (i + 1)
		coinReward := challenge.PrizeCoins / (i + 1)

		if xpReward > 0 {
			r.AddXP(p.UserID, xpReward)
		}
		if coinReward > 0 {
			r.AddCoins(p.UserID, coinReward)
		}
	}

	return nil
}

func (r *Repository) AddXP(userID, xp int) error {
	query := `UPDATE users SET total_xp = total_xp + $1 WHERE id = $2`
	_, err := r.db.Exec(query, xp, userID)
	return err
}

func (r *Repository) AddCoins(userID, coins int) error {
	query := `UPDATE users SET coins = coins + $1 WHERE id = $2`
	_, err := r.db.Exec(query, coins, userID)
	return err
}
