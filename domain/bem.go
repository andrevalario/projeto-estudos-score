package domain

type Bem struct {
	Id             uint64  `json:"id"`
	Nome           string  `json:"nome"`
	Valor          float64 `json:"valor"`
	IdProprietario uint64  `json:"proprietarioId"`
}
