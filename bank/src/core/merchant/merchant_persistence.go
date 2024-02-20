package merchant

type MerchantPersistence interface {
	Create(data *Merchant)
	Update(data *Merchant)
	GetByMerchantIdOrNil(merchantId string) *Merchant
}
