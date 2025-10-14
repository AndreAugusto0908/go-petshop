package usuario

type TipoUsuario string

const (
	TipoAdmin       TipoUsuario = "admin"
	TipoVeterinario TipoUsuario = "veterinario"
	TipoGroomer     TipoUsuario = "groomer"
	TipoAtendente   TipoUsuario = "atendente"
)
