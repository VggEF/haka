package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Login     string    `json:"login"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Group     string    `json:"group"`
	Course    int       `json:"course"`
	Faculty   string    `json:"faculty"`
	Photo     string    `json:"photo"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Telegram  string    `json:"telegram"`
	VK        string    `json:"vk"`
	GitHub    string    `json:"github"`
	TotalXP   int       `json:"total_xp"`
	Coins     int       `json:"coins"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StudentProfile struct {
	UserID        int      `json:"user_id"`
	Hobbies       []string `json:"hobbies"`
	Clubs         []string `json:"clubs"`
	About         string   `json:"about"`
	ChecklistData string   `json:"checklist_data"`
	PlannerData   string   `json:"planner_data"`
}

type TeacherProfile struct {
	UserID       int    `json:"user_id"`
	Department   string `json:"department"`
	Position     string `json:"position"`
	Degree       string `json:"degree"`
	Office       string `json:"office"`
	Experience   int    `json:"experience"`
	Achievements string `json:"achievements"`
}
