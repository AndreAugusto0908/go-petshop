package services

import (
	"context"
	"errors"
	"fmt"
	"go-petshop/internal/dtos"
	"go-petshop/internal/repositories"
)

type AuthService struct {
	repository repositories.UsuarioRepository
}

func NewAuthService(repository repositories.UsuarioRepository) AuthService {
	return AuthService{repository: repository}
}

func (s AuthService) Login(ctx context.Context, login *dtos.LoginRequest) (string, error) {
	usuario, err := s.repository.FindByEmail(ctx, login.Email)
	if err != nil {
		return "", fmt.Errorf("Erro ao realizar Login: %w", err)
	}

	if usuario == nil {
		return "", errors.New("Credenciais Invalidas")
	}

	err = usuario.VerifyPassword(login.Senha)
	if err != nil {
		return "", fmt.Errorf("Credenciais Invalidas")
	}

	tokenString, err := generateToken(login.Email)
	if err != nil {
		return "", fmt.Errorf("Erro ao generar token: %w", err)
	}

	return tokenString, nil

}
