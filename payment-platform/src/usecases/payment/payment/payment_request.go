package payment

import (
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/exception"
)

type PaymentRequest struct {
	CustomerCode string  `json:"customer_code"`
	MerchantCode string  `json:"merchant_code"`
	Amount       float64 `json:"amount"`
	CardToken    string  `json:"card_token"`
}

func (request *PaymentRequest) Validate() error {
	if request.CustomerCode == "" {
		return exception.BadRequestError(string(exception.CUSTOMER_CODE_IS_REQUIRED))
	}
	if request.MerchantCode == "" {
		return exception.BadRequestError(string(exception.MERCHANT_CODE_IS_REQUIRED))
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
