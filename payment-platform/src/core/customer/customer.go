package customer

type Customer struct {
	Id         int32  `gorm:"primaryKey" json:"id"`
	Code       string `json:"code"`
	NationalId string `json:"national_id"`
	Name       string `json:"name"`
	CreatedAt  int64  `gorm:"autoCreateTime:false" json:"created_at"`
}
