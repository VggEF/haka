package challenges

import "time"

type CreateChallengeRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Type        string `json:"type"` // tasks, grades, project
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	PrizeXP     int    `json:"prize_xp"`
	PrizeCoins  int    `json:"prize_coins"`
}

type UpdateChallengeRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	PrizeXP     int    `json:"prize_xp"`
	PrizeCoins  int    `json:"prize_coins"`
	Status      string `json:"status"`
}

type ChallengeResponse struct {
	ID                int       `json:"id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Type              string    `json:"type"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	PrizeXP           int       `json:"prize_xp"`
	PrizeCoins        int       `json:"prize_coins"`
	Status            string    `json:"status"`
	ParticipantsCount int       `json:"participants_count"`
	MyScore           int       `json:"my_score,omitempty"`
	MyRank            int       `json:"my_rank,omitempty"`
	CreatedBy         int       `json:"created_by"`
	CreatedAt         time.Time `json:"created_at"`
}

type JoinChallengeRequest struct {
	ChallengeID int `json:"challenge_id" binding:"required"`
}

type UpdateScoreRequest struct {
	UserID int `json:"user_id" binding:"required"`
	Score  int `json:"score" binding:"required"`
}

type ParticipantResponse struct {
	UserID   int       `json:"user_id"`
	UserName string    `json:"user_name"`
	Score    int       `json:"score"`
	Rank     int       `json:"rank"`
	JoinedAt time.Time `json:"joined_at"`
}

type LeaderboardResponse struct {
	ChallengeID    int                   `json:"challenge_id"`
	ChallengeTitle string                `json:"challenge_title"`
	Participants   []ParticipantResponse `json:"participants"`
	TotalCount     int                   `json:"total_count"`
}

type GetChallengesQuery struct {
	Type   string `form:"type"`
	Status string `form:"status"` // active, upcoming, ended
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}
