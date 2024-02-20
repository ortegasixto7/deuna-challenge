package merchant

type Merchant struct {
	Id         int32   `gorm:"primaryKey" json:"id"`
	MerchantId string  `json:"merchant_id"`
	Name       string  `json:"name"`
	Balance    float64 `json:"balance"`
	CreatedAt  int64   `gorm:"autoCreateTime:false" json:"created_at"`
}
