package payment

type Payment struct {
	Id           int32   `gorm:"primaryKey" json:"id"`
	Code         string  `json:"code"`
	CustomerId   string  `json:"customer_id"`
	CustomerName string  `json:"customer_name"`
	MerchantId   string  `json:"merchant_id"`
	MerchantName string  `json:"merchant_name"`
	Amount       float64 `json:"amount"`
	IsRefunded   bool    `json:"is_refunded"`
	ExternalCode string  `json:"external_code"`
	CreatedAt    int64   `gorm:"autoCreateTime:false" json:"created_at"`
}
