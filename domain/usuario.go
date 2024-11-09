package domain

type Usuario struct {
	Id          uint64 `json:"id"`
	Nome        string `json:"nome"`
	Email       string `json:"email"`
	Senha       string `json:"senha"`
	TipoUsuario uint64 `json:"tipoUsuario"`
}

type UsuarioResponse struct {
	Usuarios []Usuario `json:"usuarios"`
}
type LoginRequest struct {
	ID    uint64 `json:"id"`
	Senha string `json:"senha"`
}
