package signup

import (
	"fmt"
	"time"

	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/customer"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/exception"
)

func Execute(request *SignUpRequest, customerPersistence customer.CustomerPersistence) (*customer.Customer, error) {
	requestErrorCode := request.Validate()
	if requestErrorCode != nil {
		return nil, requestErrorCode
	}

	customerResult := customerPersistence.GetByNationalIdOrNil(request.NationalId)
	if customerResult != nil {
		return nil, exception.BadRequestError(string(exception.UNAVAILABLE_NATIONAL_ID))
	}

	var customer customer.Customer
	customer.Id = 0
	customer.NationalId = request.NationalId
	customer.Name = request.Name
	customer.CreatedAt = time.Now().UnixMilli()
	customer.Code = fmt.Sprintf("CUS_%d", customer.CreatedAt)

	customerPersistence.Create(&customer)

	return &customer, nil
}
