package transaction

type TransactionPersistence interface {
	Create(data *Transaction)
	Update(data *Transaction)
	GetByCodeOrNil(code string) *Transaction
}
