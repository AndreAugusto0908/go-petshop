package usuario

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	Id          string      `json:"id"`
	Nome        string      `json:"nome"`
	Email       string      `json:"email"`
	Senha       string      `json:"senha"`
	Tipo        TipoUsuario `json:"tipo"`
	CRMV        *string     `json:"crmv"`
	Telefone    *string     `json:"telefone"`
	Ativo       bool        `json:"ativo"`
	UltimoLogin *time.Time  `json:"ultimoLogin"`
	CreatedAt   *time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time  `json:"updatedAt"`
}

func (u *Usuario) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Senha = string(hash)
	return nil
}

func (u *Usuario) VerifyPassword(senhaParaVerificar string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(senhaParaVerificar))
}
