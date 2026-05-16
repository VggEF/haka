package skilltree

import "time"

type CreateSkillRequest struct {
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description"`
	XPCost          int    `json:"xp_cost" binding:"required"`
	Icon            string `json:"icon"`
	Category        string `json:"category" binding:"required"` // programming, softskills, career
	ParentSkillID   *int   `json:"parent_skill_id"`
	RequiredSkillID *int   `json:"required_skill_id"`
}

type UpdateSkillRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	XPCost          int    `json:"xp_cost"`
	Icon            string `json:"icon"`
	Category        string `json:"category"`
	ParentSkillID   *int   `json:"parent_skill_id"`
	RequiredSkillID *int   `json:"required_skill_id"`
}

type SkillResponse struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	XPCost          int        `json:"xp_cost"`
	Icon            string     `json:"icon"`
	Category        string     `json:"category"`
	ParentSkillID   *int       `json:"parent_skill_id"`
	RequiredSkillID *int       `json:"required_skill_id"`
	IsUnlocked      bool       `json:"is_unlocked"`
	CanUnlock       bool       `json:"can_unlock"`
	UnlockedAt      *time.Time `json:"unlocked_at,omitempty"`
}

type UnlockSkillRequest struct {
	SkillID int `json:"skill_id" binding:"required"`
}

type GetSkillsQuery struct {
	Category string `form:"category"`
	Unlocked bool   `form:"unlocked"`
}
