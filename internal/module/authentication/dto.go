package authentication

import "github.com/adityarifqyfauzan/go-boilerplate/internal/model"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Name                 string `json:"name" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
}

type LoginResponse struct {
	Token        string     `json:"token"`
	RefreshToken string     `json:"refresh_token"`
	User         MeResponse `json:"user"`
}

type RegisterResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}

type MeResponse struct {
	ID    int           `json:"id"`
	Name  string        `json:"name"`
	Email string        `json:"email"`
	Roles []*model.Role `json:"roles"`
}
