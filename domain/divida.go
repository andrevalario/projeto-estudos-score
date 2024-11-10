package domain

type Divida struct {
	Id        uint64  `json:"id"`
	IdUsuario uint64  `json:"id_usuario"`
	Valor     float64 `json:"valor"`
}

type DividaResponse struct {
	Dividas []Divida `json:"dividas"`
}
