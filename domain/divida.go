package domain

type Divida struct {
	Id             uint64  `json:"id"`
	UsuarioID      uint64  `json:"usuario_id"`
	Valor          float64 `json:"valor"`
	DataVencimento string  `json:"data_vencimento"`
	Status         string  `json:"status"`
}
