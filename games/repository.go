package games

import (
	"database/sql"
	"fmt"
	"strings"
	"student-app/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// ========== Игры ==========
func (r *Repository) CreateGame(game *models.MiniGame) error {
	query := `
        INSERT INTO mini_games (title, description, icon, xp_reward, coin_reward, game_data)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at
    `
	return r.db.QueryRow(query, game.Title, game.Description, game.Icon,
		game.XPReward, game.CoinReward, game.GameData).Scan(&game.ID, &game.CreatedAt)
}

func (r *Repository) UpdateGame(id int, updates map[string]interface{}) error {
	setParts := make([]string, 0, len(updates))
	args := make([]interface{}, 0, len(updates)+1)
	i := 1

	for key, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE mini_games SET %s WHERE id = $%d", strings.Join(setParts, ", "), i)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Repository) DeleteGame(id int) error {
	query := `DELETE FROM mini_games WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) FindGameByID(id int) (*models.MiniGame, error) {
	var game models.MiniGame
	query := `
        SELECT id, title, description, icon, xp_reward, coin_reward, game_data, created_at
        FROM mini_games
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(&game.ID, &game.Title, &game.Description,
		&game.Icon, &game.XPReward, &game.CoinReward, &game.GameData, &game.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (r *Repository) GetAllGames(isActive bool, limit, offset int) ([]models.MiniGame, error) {
	query := `
        SELECT id, title, description, icon, xp_reward, coin_reward, created_at
        FROM mini_games
        WHERE 1=1
    `
	args := []interface{}{}
	i := 1

	if isActive {
		query += " AND is_active = true"
	}

	if limit == 0 {
		limit = 50
	}
	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.MiniGame
	for rows.Next() {
		var game models.MiniGame
		err := rows.Scan(&game.ID, &game.Title, &game.Description, &game.Icon,
			&game.XPReward, &game.CoinReward, &game.CreatedAt)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}

func (r *Repository) GetGameStats(gameID int) (playsCount int, avgScore float64, err error) {
	query := `
        SELECT COUNT(*), COALESCE(AVG(score), 0)
        FROM game_results
        WHERE game_id = $1
    `
	err = r.db.QueryRow(query, gameID).Scan(&playsCount, &avgScore)
	return
}

// ========== Результаты игр ==========
func (r *Repository) SaveGameResult(result *models.GameResult) error {
	query := `
        INSERT INTO game_results (user_id, game_id, score, xp_earned, coins_earned)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	return r.db.QueryRow(query, result.UserID, result.GameID, result.Score,
		result.XPEarned, result.CoinsEarned).Scan(&result.ID)
}

func (r *Repository) GetUserGameResults(userID, gameID int, limit, offset int) ([]models.GameResult, error) {
	query := `
        SELECT id, user_id, game_id, score, xp_earned, coins_earned, played_at
        FROM game_results
        WHERE user_id = $1
    `
	args := []interface{}{userID}
	i := 2

	if gameID > 0 {
		query += fmt.Sprintf(" AND game_id = $%d", i)
		args = append(args, gameID)
		i++
	}

	if limit == 0 {
		limit = 20
	}
	query += fmt.Sprintf(" ORDER BY played_at DESC LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.GameResult
	for rows.Next() {
		var r models.GameResult
		err := rows.Scan(&r.ID, &r.UserID, &r.GameID, &r.Score,
			&r.XPEarned, &r.CoinsEarned, &r.PlayedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return results, nil
}

func (r *Repository) GetUserBestScore(userID, gameID int) (int, error) {
	var bestScore int
	query := `SELECT COALESCE(MAX(score), 0) FROM game_results WHERE user_id = $1 AND game_id = $2`
	err := r.db.QueryRow(query, userID, gameID).Scan(&bestScore)
	return bestScore, err
}

func (r *Repository) GetGameLeaderboard(gameID int, limit int) ([]GameLeaderboardEntry, error) {
	query := `
        SELECT gr.user_id, u.name, MAX(gr.score) as best_score, MAX(gr.played_at) as last_played
        FROM game_results gr
        JOIN users u ON gr.user_id = u.id
        WHERE gr.game_id = $1
        GROUP BY gr.user_id, u.name
        ORDER BY best_score DESC, last_played ASC
        LIMIT $2
    `
	rows, err := r.db.Query(query, gameID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []GameLeaderboardEntry
	for rows.Next() {
		var entry GameLeaderboardEntry
		err := rows.Scan(&entry.UserID, &entry.UserName, &entry.Score, &entry.PlayedAt)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

// ========== Вспомогательные ==========
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
