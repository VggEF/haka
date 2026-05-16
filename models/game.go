package models

import "time"

type MiniGame struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	XPReward    int       `json:"xp_reward"`
	CoinReward  int       `json:"coin_reward"`
	GameData    string    `json:"game_data"`
	CreatedAt   time.Time `json:"created_at"`
}

type GameResult struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	GameID      int       `json:"game_id"`
	Score       int       `json:"score"`
	XPEarned    int       `json:"xp_earned"`
	CoinsEarned int       `json:"coins_earned"`
	PlayedAt    time.Time `json:"played_at"`
}
