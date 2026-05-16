package users

type CreateUserRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role"`
	Group    string `json:"group"`
	Course   int    `json:"course"`
	Faculty  string `json:"faculty"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Telegram string `json:"telegram"`
	VK       string `json:"vk"`
	GitHub   string `json:"github"`
	Photo    string `json:"photo"`
	Group    string `json:"group"`
	Course   int    `json:"course"`
	Faculty  string `json:"faculty"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Group    string `json:"group"`
	Course   int    `json:"course"`
	Faculty  string `json:"faculty"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Telegram string `json:"telegram"`
	VK       string `json:"vk"`
	GitHub   string `json:"github"`
	Photo    string `json:"photo"`
	TotalXP  int    `json:"total_xp"`
	Coins    int    `json:"coins"`
}

type UpdateStudentProfileRequest struct {
	Hobbies []string `json:"hobbies"`
	Clubs   []string `json:"clubs"`
	About   string   `json:"about"`
}

type UpdateTeacherProfileRequest struct {
	Department string `json:"department"`
	Position   string `json:"position"`
	Degree     string `json:"degree"`
	Office     string `json:"office"`
	Experience int    `json:"experience"`
}

type StudentProfileResponse struct {
	Hobbies []string `json:"hobbies"`
	Clubs   []string `json:"clubs"`
	About   string   `json:"about"`
}

type TeacherProfileResponse struct {
	Department string `json:"department"`
	Position   string `json:"position"`
	Degree     string `json:"degree"`
	Office     string `json:"office"`
	Experience int    `json:"experience"`
}
