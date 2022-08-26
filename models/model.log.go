package models

type ModelLog struct {
	ID          int     `gorm:"primaryKey" json:"id"`
	SystemName  string  `json:"system_name"`
	ActionType  string  `json:"action_type"`
	TblName     string  `json:"table_name"`
	Url         string  `json:"url"`
	Request     string  `json:"request"`
	Response    string  `json:"response"`
	StatusCode  int     `json:"status_code"`
	Duration    float32 `json:"duration"`
	AppId       int     `json:"app_id"`
	CreatedTime int     `gorm:"autoCreateTime"`
	UpdatedTime int     `gorm:"autoUpdateTime"`
}

func (m *ModelLog) TableName() string {
	return "open_api_request_log"
}
