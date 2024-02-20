package refund

import (
	"github.com/ortegasixto7/deuna-challenge/bank/src/core/customer"
	"github.com/ortegasixto7/deuna-challenge/bank/src/core/merchant"
	"github.com/ortegasixto7/deuna-challenge/bank/src/core/transaction"
	"github.com/ortegasixto7/deuna-challenge/bank/src/exception"
)

func Execute(
	request *RefundRequest,
	customerPersistence customer.CustomerPersistence,
	merchantPersistence merchant.MerchantPersistence,
	transactionPersistence transaction.TransactionPersistence) (*transaction.Transaction, error) {
	requestErrorCode := request.Validate()
	if requestErrorCode != nil {
		return nil, requestErrorCode
	}

	transaction := transactionPersistence.GetByCodeOrNil(request.TransactionCode)
	if transaction == nil {
		return nil, exception.NotFoundError(string(exception.TRANSACTION_NOT_FOUND))
	}

	if transaction.IsRefunded {
		return nil, exception.BadRequestError(string(exception.TRANSACTION_IS_ALREADY_REFUNDED))
	}

	customer := customerPersistence.GetByNationalIdOrNil(transaction.CustomerId)
	if customer == nil {
		return nil, exception.NotFoundError(string(exception.CUSTOMER_NOT_FOUND))
	}

	merchant := merchantPersistence.GetByMerchantIdOrNil(transaction.MerchantId)
	if merchant == nil {
		return nil, exception.NotFoundError(string(exception.MERCHANT_NOT_FOUND))
	}

	transaction.IsRefunded = true

	customer.Balance += transaction.Amount
	merchant.Balance -= transaction.Amount

	customerPersistence.Update(customer)
	merchantPersistence.Update(merchant)
	transactionPersistence.Update(transaction)

	return transaction, nil
}
