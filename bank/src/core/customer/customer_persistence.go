package customer

type CustomerPersistence interface {
	Create(data *Customer)
	Update(data *Customer)
	GetByNationalIdOrNil(nationalId string) *Customer
}
