package exception

type BadRequestException struct {
	ErrorMessage string
}

func (badRequestException *BadRequestException) Error() string {
	return badRequestException.ErrorMessage
}

func BadRequestError(errorMessage string) error {
	return &BadRequestException{ErrorMessage: errorMessage}
}

type BadRequest string

const (
	NATIONAL_ID_IS_REQUIRED     BadRequest = "NATIONAL_ID_IS_REQUIRED"
	MERCHANT_ID_IS_REQUIRED     BadRequest = "MERCHANT_ID_IS_REQUIRED"
	CUSTOMER_ID_IS_REQUIRED     BadRequest = "CUSTOMER_ID_IS_REQUIRED"
	MERCHANT_CODE_IS_REQUIRED   BadRequest = "MERCHANT_CODE_IS_REQUIRED"
	CUSTOMER_CODE_IS_REQUIRED   BadRequest = "CUSTOMER_CODE_IS_REQUIRED"
	PAYMENT_CODE_IS_REQUIRED    BadRequest = "PAYMENT_CODE_IS_REQUIRED"
	CARD_TOKEN_IS_REQUIRED      BadRequest = "CARD_TOKEN_IS_REQUIRED"
	NAME_IS_REQUIRED            BadRequest = "NAME_IS_REQUIRED"
	INVALID_AMOUNT              BadRequest = "INVALID_AMOUNT"
	INSUFFICIENT_FUNDS          BadRequest = "INSUFFICIENT_FUNDS"
	INVALID_TRANSACTION         BadRequest = "INVALID_TRANSACTION"
	INVALID_CARD_TOKEN          BadRequest = "INVALID_CARD_TOKEN"
	PAYMENT_IS_ALREADY_REFUNDED BadRequest = "PAYMENT_IS_ALREADY_REFUNDED"
	PAYMENT_FAILED              BadRequest = "PAYMENT_FAILED"
	PAYMENT_DECLINED            BadRequest = "PAYMENT_DECLINED"
	REFUND_FAILED               BadRequest = "REFUND_FAILED"

	UNAVAILABLE_NATIONAL_ID BadRequest = "UNAVAILABLE_NATIONAL_ID"
	UNAVAILABLE_MERCHANT_ID BadRequest = "UNAVAILABLE_MERCHANT_ID"
)
