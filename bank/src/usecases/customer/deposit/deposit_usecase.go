package deposit

import (
	"github.com/ortegasixto7/deuna-challenge/bank/src/core/customer"
	"github.com/ortegasixto7/deuna-challenge/bank/src/exception"
)

func Execute(request *DepositRequest, customerPersistence customer.CustomerPersistence) (*customer.Customer, error) {
	requestErrorCode := request.Validate()
	if requestErrorCode != nil {
		return nil, requestErrorCode
	}

	customer := customerPersistence.GetByNationalIdOrNil(request.NationalId)
	if customer == nil {
		return nil, exception.NotFoundError(string(exception.CUSTOMER_NOT_FOUND))
	}
	customer.Balance += request.Amount
	customerPersistence.Update(customer)

	return customer, nil
}
