package models

import "time"

type Skill struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	XPCost          int       `json:"xp_cost"`
	Icon            string    `json:"icon"`
	Category        string    `json:"category"`
	ParentSkillID   *int      `json:"parent_skill_id"`
	RequiredSkillID *int      `json:"required_skill_id"`
	CreatedAt       time.Time `json:"created_at"`
}

type UserSkill struct {
	UserID     int       `json:"user_id"`
	SkillID    int       `json:"skill_id"`
	UnlockedAt time.Time `json:"unlocked_at"`
}
