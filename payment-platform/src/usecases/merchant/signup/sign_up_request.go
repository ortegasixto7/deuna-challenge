package signup

import (
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/exception"
)

type SignUpRequest struct {
	MerchantId string `json:"merchant_id"`
	Name       string `json:"name"`
}

func (request *SignUpRequest) Validate() error {
	if request.MerchantId == "" {
		return exception.BadRequestError(string(exception.MERCHANT_ID_IS_REQUIRED))
	}
	if request.Name == "" {
		return exception.BadRequestError(string(exception.NAME_IS_REQUIRED))
	}
	return nil
}
