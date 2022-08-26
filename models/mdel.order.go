package models

type ModelOrder struct {
	ID                int    `gorm:"primaryKey" json:"id"`
	RefCustomerId     string `json:"ref_customer_id"`
	RefTransactionId  string `json:"ref_transaction_id"`
	RefTransactionIdx int    `json:"refTransactionIdx"`
	Process           int    `json:"process"`
	UrboxStatus       int    `json:"urbox_status"`
	CreatedAt         int    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         int    `gorm:"autoUpdateTime" json:"updated_at"`
	Status            int    `json:"status"`
}

func (m *ModelOrder) TableName() string {
	return "order"
}
