package models

type ModelGiftCode struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	GiftDetailId int    `json:"gift_detail_id"`
	Link         string `json:"link"`
	ValidTime    int    `json:"valid_time"`
	ExpiredTime  int    `json:"expired_time"`
	Process      int    `json:"process"`
	Status       int    `json:"status"`
}

func (m *ModelGiftCode) TableName() string {
	return "gift_code"
}
