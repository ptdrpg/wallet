package entity

type Transaction struct {
	UID    string  `gorm:"primary_key" json:"uid"`
	Type   string  `json:"type"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}
