package avaliacao

type StatusAvaliacao string

const (
	StatusPendente    StatusAvaliacao = "pendente"
	StatusEmAndamento StatusAvaliacao = "em_andamento"
	StatusConcluida   StatusAvaliacao = "concluida"
	StatusCancelada   StatusAvaliacao = "cancelada"
)
