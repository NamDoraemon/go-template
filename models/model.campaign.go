package models

type ModelCampaign struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	AppId       int    `json:"app_id"`
	AppSecret   string `json:"app_secret"`
	PrivateKey  string `json:"private_key"`
	MerchantKey string `json:"merchant_key"`
	Status      int    `json:"status"`
}

func (m *ModelCampaign) TableName() string {
	return "campaign"
}
