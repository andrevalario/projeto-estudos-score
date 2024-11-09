package domain

type ApiError struct {
	Id       uint64 `json:"id"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Status   int    `json:"status"`
	Original error  `json:"-"`
}

func (a *ApiError) Error() string {
	return a.Detail
}
