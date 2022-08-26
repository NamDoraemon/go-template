package models

type ModelOrderDetail struct {
	ID                 int    `gorm:"primaryKey" json:"id"`
	OrderId            int    `json:"order_id"`
	VoucherId          int    `json:"voucher_id"`
	VoucherDetailId    int    `json:"voucher_detail_id"`
	VoucherName        string `json:"voucher_name"`
	Point              int    `json:"point"`
	Price              int    `json:"price"`
	AppId              int    `json:"app_id"`
	UrboxTransactionId string `json:"urbox_transaction_id"`
	Code               string `json:"code"`
	Link               string `json:"link"`
	Process            int    `json:"process"`
	UrboxStatus        int    `json:"urbox_status"`
	CreatedAt          int    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          int    `gorm:"autoUpdateTime" json:"updated_at"`
}

func (m *ModelOrderDetail) TableName() string {
	return "order_detail"
}
