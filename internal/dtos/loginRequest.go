package dtos

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	Senha string `json:"senha" validate:"required,min=6"`
}
