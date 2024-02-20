package exception

type NotFoundException struct {
	ErrorMessage string
}

func (notFoundException *NotFoundException) Error() string {
	return notFoundException.ErrorMessage
}

func NotFoundError(errorMessage string) error {
	return &NotFoundException{ErrorMessage: errorMessage}
}

type NotFound string

const (
	CUSTOMER_NOT_FOUND    NotFound = "CUSTOMER_NOT_FOUND"
	MERCHANT_NOT_FOUND    NotFound = "MERCHANT_NOT_FOUND"
	TRANSACTION_NOT_FOUND NotFound = "TRANSACTION_NOT_FOUND"
)
