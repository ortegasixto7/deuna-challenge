package transfer

import (
	"fmt"
	"time"

	"github.com/ortegasixto7/deuna-challenge/bank/src/core/customer"
	"github.com/ortegasixto7/deuna-challenge/bank/src/core/merchant"
	"github.com/ortegasixto7/deuna-challenge/bank/src/core/transaction"
	"github.com/ortegasixto7/deuna-challenge/bank/src/exception"
)

func Execute(
	request *TransferRequest,
	customerPersistence customer.CustomerPersistence,
	merchantPersistence merchant.MerchantPersistence,
	transactionPersistence transaction.TransactionPersistence) (*transaction.Transaction, error) {
	requestErrorCode := request.Validate()
	if requestErrorCode != nil {
		return nil, requestErrorCode
	}

	customer := customerPersistence.GetByNationalIdOrNil(request.CustomerId)
	if customer == nil {
		return nil, exception.NotFoundError(string(exception.CUSTOMER_NOT_FOUND))
	}

	if customer.Balance < request.Amount {
		return nil, exception.BadRequestError(string(exception.INSUFFICIENT_FUNDS))
	}

	merchant := merchantPersistence.GetByMerchantIdOrNil(request.MerchantId)
	if merchant == nil {
		return nil, exception.NotFoundError(string(exception.MERCHANT_NOT_FOUND))
	}

	var transaction transaction.Transaction
	transaction.Id = 0
	transaction.Code = fmt.Sprintf("TX_%d", time.Now().UnixMilli())
	transaction.CustomerId = customer.NationalId
	transaction.CustomerName = customer.Name
	transaction.MerchantId = merchant.MerchantId
	transaction.MerchantName = merchant.Name
	transaction.Amount = request.Amount
	transaction.IsRefunded = false
	transaction.CreatedAt = time.Now().UnixMilli()

	customer.Balance -= request.Amount
	merchant.Balance += request.Amount

	customerPersistence.Update(customer)
	merchantPersistence.Update(merchant)
	transactionPersistence.Create(&transaction)

	return &transaction, nil
}
