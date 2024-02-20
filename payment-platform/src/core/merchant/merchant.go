package merchant

type Merchant struct {
	Id         int32  `gorm:"primaryKey" json:"id"`
	Code       string `json:"code"`
	MerchantId string `json:"merchant_id"`
	Name       string `json:"name"`
	CreatedAt  int64  `gorm:"autoCreateTime:false" json:"created_at"`
}
