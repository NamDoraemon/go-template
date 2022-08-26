package models

type ModelVoucherDetail struct {
	ID                    int    `gorm:"primaryKey" json:"id"`
	AppId                 int    `json:"app_id"`
	VoucherId             int    `json:"voucher_id"`
	RefId                 int    `json:"ref_id"`
	NumberCode            int    `json:"number_code"`
	SalesLimit            int    `json:"sales_limit"`
	ThmbnailImage         string `json:"thmbnail_image"`
	Image                 string `json:"image"`
	MyViettelDisplayValue string `json:"my_viettel_display_value"`
	StartTime             int    `json:"start_time"`
	EndTime               int    `json:"end_time"`
	Price                 int    `json:"price"`
	UrboxGiftId           int    `json:"urbox_gift_id"`
	UrboxGiftDetailId     int    `json:"urbox_gift_detail_id"`
	AutoOff               int    `json:"auto_off"`
	UrboxStatus           int    `json:"urbox_status"`
	Status                int    `json:"status"`
	RewardPoint           int    `json:"reward_point"`
	MaxRewardPoint        int    `json:"max_reward_point"`
	RewardPointPercent    int    `json:"reward_point_percent"`
	ExchangePoint         int    `json:"exchange_point"`
	SalePercent           int    `json:"sale_percent"`
	RequiredStamps        int    `json:"required_stamps"`
	MinimumVtPointToJoin  int    `json:"minimum_vt_point_to_join"`
	CreatedTime           int    `gorm:"autoCreateTime" json:"created_time"`
	UpdatedTime           int    `gorm:"autoUpdateTime" json:"updated_time"`
}

func (m *ModelVoucherDetail) TableName() string {
	return "voucher_detail"
}
