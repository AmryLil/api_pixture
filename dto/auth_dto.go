package dto

type RegisterRequest struct {
	Fullname        string `json:"fullname" binding:"required"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"Name"`
	Token string `json:"token"`
}
