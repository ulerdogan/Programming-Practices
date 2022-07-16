package domain

type Balance struct {
	Balance string `json:"balance"`
	Nonce uint `json:"nonce"`
	Network string `json:"network"`
}