package cao

import (
	"time"
)

type Cao struct {
	ID          string    `json:"id" db:"id"`
	Nome        string    `json:"nome" db:"nome"`
	RacaId      string    `json:"raca_id" db:"raca"`
	Idade       *int      `json:"idade,omitempty" db:"idade"`
	Peso        *float64  `json:"peso,omitempty" db:"peso"`
	NomeDono    string    `json:"nome_dono" db:"nome_dono"`
	ContatoDono string    `json:"contato_dono" db:"contato_dono"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
