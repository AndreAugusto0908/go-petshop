package raca

import "time"

type Raca struct {
	Id        string     `json:"id"`
	Nome      string     `json:"nome"`
	Raca      string     `json:"raca"`
	Descricao string     `json:"descricao"`
	Createdat *time.Time `json:"createdat"`
	Updatedat *time.Time `json:"updatedat"`
}
