package entities

type RequestHoldCode struct {
	CustomerId    string `json:"customerId" binding:"required"`
	TransactionId string `json:"transactionId" binding:"required"`
	Promotion     int    `json:"promotion" binding:"required,gte=1"`
}

type RequestGetCode struct {
	RequestHoldCode
	PointType int `json:"pointType" binding:"required,default=123"`
}
