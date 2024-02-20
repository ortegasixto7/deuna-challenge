package getbycode

import (
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/payment"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/exception"
)

func Execute(requestCode string, paymentPersistence payment.PaymentPersistence) (*payment.Payment, error) {
	payment := paymentPersistence.GetByCodeOrNil(requestCode)
	if payment == nil {
		return nil, exception.NotFoundError(string(exception.PAYMENT_NOT_FOUND))
	}
	return payment, nil
}
