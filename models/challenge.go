package models

import "time"

type Challenge struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	PrizeXP     int       `json:"prize_xp"`
	PrizeCoins  int       `json:"prize_coins"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type ChallengeParticipant struct {
	ChallengeID int       `json:"challenge_id"`
	UserID      int       `json:"user_id"`
	Score       int       `json:"score"`
	JoinedAt    time.Time `json:"joined_at"`
}
