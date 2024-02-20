package payment

type PaymentPersistence interface {
	Create(data *Payment)
	Update(data *Payment)
	GetByCodeOrNil(code string) *Payment
}
