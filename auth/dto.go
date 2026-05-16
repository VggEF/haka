package auth

type LoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role"`
	Group    string `json:"group"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	ID      int    `json:"id"`
	Login   string `json:"login"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Group   string `json:"group"`
	TotalXP int    `json:"total_xp"`
	Coins   int    `json:"coins"`
}
