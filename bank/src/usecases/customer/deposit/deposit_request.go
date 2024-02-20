package deposit

import (
	"github.com/ortegasixto7/deuna-challenge/bank/src/exception"
)

type DepositRequest struct {
	NationalId string  `json:"national_id"`
	Amount     float64 `json:"amount"`
}

func (request *DepositRequest) Validate() error {
	if request.NationalId == "" {
		return exception.BadRequestError(string(exception.NATIONAL_ID_IS_REQUIRED))
	}
	if request.Amount <= 0 {
		return exception.BadRequestError(string(exception.INVALID_AMOUNT))
	}
	return nil
}
