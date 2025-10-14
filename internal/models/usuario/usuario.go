package usuario

import (
	"time"
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
