package services

import (
	"context"
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
	usuario, err := s.repository.Create(ctx, entity)

	if !entity.Ativo {
		entity.Ativo = true
	}

	if err != nil {
		return u.Usuario{}, err
	}

	entity.Id = usuario.Id
	return *usuario, nil
}
