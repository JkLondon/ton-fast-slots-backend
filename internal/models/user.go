package models

type RegisterUserArgs struct {
	TGID          int64
	WalletAddress string
	WalletSeed    string
}

type GetSelfInfoResponse struct {
	TGID          int64   `db:"tg_id" json:"TGID"`
	WalletAddress string  `db:"wallet_address" json:"walletAddress"`
	Balance       float64 `json:"balance"`
}
