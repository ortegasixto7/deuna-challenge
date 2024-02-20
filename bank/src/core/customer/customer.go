package customer

type Customer struct {
	Id         int32   `gorm:"primaryKey" json:"id"`
	NationalId string  `json:"national_id"`
	Name       string  `json:"name"`
	Balance    float64 `json:"balance"`
	CreatedAt  int64   `gorm:"autoCreateTime:false" json:"created_at"`
}
