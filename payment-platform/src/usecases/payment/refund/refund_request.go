package refund

import (
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/exception"
)

type RefundRequest struct {
	PaymentCode string `json:"payment_code"`
}

func (request *RefundRequest) Validate() error {
	if request.PaymentCode == "" {
		return exception.BadRequestError(string(exception.PAYMENT_CODE_IS_REQUIRED))
	}

	return nil
}
