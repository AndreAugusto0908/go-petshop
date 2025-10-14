package avaliacao

import "time"

type Avaliacao struct {
	ID              string          `json:"id" db:"id"`
	CaoID           string          `json:"cao_id" db:"cao_id"`
	UsuarioID       string          `json:"usuario_id" db:"usuario_id"`
	Tipo            TipoAvaliacao   `json:"tipo" db:"tipo"`
	Status          StatusAvaliacao `json:"status" db:"status"`
	Nota            *int            `json:"nota,omitempty" db:"nota"`
	Titulo          string          `json:"titulo" db:"titulo"`
	Descricao       *string         `json:"descricao,omitempty" db:"descricao"`
	Observacoes     *string         `json:"observacoes,omitempty" db:"observacoes"`
	Recomendacoes   *string         `json:"recomendacoes,omitempty" db:"recomendacoes"`
	ProximaConsulta *time.Time      `json:"proxima_consulta,omitempty" db:"proxima_consulta"`
	DataAvaliacao   time.Time       `json:"data_avaliacao" db:"data_avaliacao"`
	CreatedAt       time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at" db:"updated_at"`
}
