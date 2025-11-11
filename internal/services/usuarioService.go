package services

import (
	"context"
	"fmt"
	u "go-petshop/internal/models/usuario"
	"go-petshop/internal/repositories"
)

type UsuarioService struct {
	repository repositories.UsuarioRepository
}

func NewUsuarioService(repository repositories.UsuarioRepository) UsuarioService {
	return UsuarioService{repository: repository}
}

func (s UsuarioService) Create(ctx context.Context, entity *u.Usuario) (u.Usuario, error) {

	if !entity.Ativo {
		entity.Ativo = true
	}

	err := entity.EncryptPassword()
	if err != nil {
		return u.Usuario{}, fmt.Errorf("erro ao encriptar senha: %w", err)
	}

	usuario, err := s.repository.Create(ctx, entity)

	if err != nil {
		return u.Usuario{}, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	entity.Id = usuario.Id
	return *usuario, nil
}
