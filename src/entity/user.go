package entity

type Users struct {
	ID       string
	Name     string
	Email    string
	Password string
	Address  string
}

type UserRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address" validate:"required"`
}

type UserResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
