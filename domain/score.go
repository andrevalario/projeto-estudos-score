package domain

type ScoreRequest struct {
	Name        string  `json:"name"`
	Income      float64 `json:"income"`
	Debt        float64 `json:"debt"`
	CreditLimit float64 `json:"credit_limit"`
}
