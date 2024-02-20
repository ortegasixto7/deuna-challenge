package refund

import (
	"github.com/ortegasixto7/deuna-challenge/bank/src/exception"
)

type RefundRequest struct {
	TransactionCode string `json:"transaction_code"`
}

func (request *RefundRequest) Validate() error {
	if request.TransactionCode == "" {
		return exception.BadRequestError(string(exception.TRANSACTION_CODE_IS_REQUIRED))
	}

	return nil
}
