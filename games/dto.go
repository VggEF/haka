package games

import "time"

type CreateGameRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	XPReward    int    `json:"xp_reward"`
	CoinReward  int    `json:"coin_reward"`
	GameData    string `json:"game_data"` // JSON с данными игры
}

type UpdateGameRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	XPReward    int    `json:"xp_reward"`
	CoinReward  int    `json:"coin_reward"`
	GameData    string `json:"game_data"`
	IsActive    bool   `json:"is_active"`
}

type GameResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	XPReward    int       `json:"xp_reward"`
	CoinReward  int       `json:"coin_reward"`
	GameData    string    `json:"game_data"`
	IsActive    bool      `json:"is_active"`
	PlaysCount  int       `json:"plays_count"`
	AvgScore    float64   `json:"avg_score"`
	CreatedAt   time.Time `json:"created_at"`
}

type SubmitGameResultRequest struct {
	GameID int `json:"game_id" binding:"required"`
	Score  int `json:"score" binding:"required"`
}

type GameResultResponse struct {
	ID          int       `json:"id"`
	GameID      int       `json:"game_id"`
	GameTitle   string    `json:"game_title"`
	Score       int       `json:"score"`
	XPEarned    int       `json:"xp_earned"`
	CoinsEarned int       `json:"coins_earned"`
	PlayedAt    time.Time `json:"played_at"`
}

type GetGamesQuery struct {
	IsActive bool `form:"is_active"`
	Limit    int  `form:"limit"`
	Offset   int  `form:"offset"`
}

type GameLeaderboardEntry struct {
	UserID   int       `json:"user_id"`
	UserName string    `json:"user_name"`
	Score    int       `json:"score"`
	PlayedAt time.Time `json:"played_at"`
}
