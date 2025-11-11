package repositories

import (
	"context"
	"database/sql"
	u "go-petshop/internal/models/usuario"
)

type UsuarioRepository struct {
	db *sql.DB
}

func NewUsuarioRepository(db *sql.DB) UsuarioRepository {
	return UsuarioRepository{db: db}
}

var _ Repository[u.Usuario] = (*UsuarioRepository)(nil)

func (r UsuarioRepository) Create(ctx context.Context, entity *u.Usuario) (*u.Usuario, error) {
	query := `INSERT INTO usuarios (nome, email, senha_hash, tipo, crmv, telefone, ativo)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)
			  RETURNING id, nome,email, senha_hash, tipo, crmv, telefone, ativo,   
		          ultimo_login, created_at, updated_at`

	newUser := &u.Usuario{}

	err := r.db.QueryRowContext(ctx, query,
		entity.Nome,
		entity.Email,
		entity.Senha,
		entity.Tipo,
		entity.CRMV,
		entity.Telefone,
		entity.Ativo,
	).Scan(
		&newUser.Id,
		&newUser.Nome,
		&newUser.Email,
		&newUser.Senha,
		&newUser.Tipo,
		&newUser.CRMV,
		&newUser.Telefone,
		&newUser.Ativo,
		&newUser.UltimoLogin,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return newUser, nil

}

func (r UsuarioRepository) FindByEmail(ctx context.Context, email string) (*u.Usuario, error) {
	query := `SELECT * FROM usuarios WHERE email = $1`

	user := &u.Usuario{}

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.Nome,
		&user.Email,
		&user.Senha,
		&user.Tipo,
		&user.CRMV,
		&user.Telefone,
		&user.Ativo,
		&user.UltimoLogin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r UsuarioRepository) FindAll(ctx context.Context) ([]*u.Usuario, error) {
	//TODO implement me
	panic("implement me")
}

func (r UsuarioRepository) Update(ctx context.Context, id string, entity *u.Usuario) error {
	//TODO implement me
	panic("implement me")
}

func (r UsuarioRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (r UsuarioRepository) FindById(ctx context.Context, id string) (*u.Usuario, error) {
	//TODO implement me
	panic("implement me")
}
