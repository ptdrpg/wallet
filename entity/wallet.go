package entity

type Wallet struct {
	UID       string  `gorm:"primary_key" json:"uid"`
	OwnerUID  string  `json:"owner_uid"`
	PublicKey string  `json:"public_key"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
}
