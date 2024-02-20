package transfer

import (
	"github.com/ortegasixto7/deuna-challenge/bank/src/exception"
)

type TransferRequest struct {
	CustomerId string  `json:"customer_id"`
	MerchantId string  `json:"merchant_id"`
	Amount     float64 `json:"amount"`
	CardToken  string  `json:"card_token"`
}

func (request *TransferRequest) Validate() error {
	if request.CustomerId == "" {
		return exception.BadRequestError(string(exception.CUSTOMER_ID_IS_REQUIRED))
	}
	if request.MerchantId == "" {
		return exception.BadRequestError(string(exception.MERCHANT_ID_IS_REQUIRED))
	}
	if request.Amount < 1 {
		return exception.BadRequestError(string(exception.INVALID_AMOUNT))
	}
	if request.CardToken == "" {
		return exception.BadRequestError(string(exception.CARD_TOKEN_IS_REQUIRED))
	}
	if request.CardToken != "tok_declined" && request.CardToken != "tok_approved" {
		return exception.BadRequestError(string(exception.INVALID_CARD_TOKEN))
	}
	if request.CardToken == "tok_declined" {
		return exception.BadRequestError(string(exception.INVALID_TRANSACTION))
	}

	return nil
}
