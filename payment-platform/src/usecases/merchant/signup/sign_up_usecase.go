package signup

import (
	"fmt"
	"time"

	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/merchant"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/exception"
)

func Execute(request *SignUpRequest, merchantPersistence merchant.MerchantPersistence) (*merchant.Merchant, error) {
	requestErrorCode := request.Validate()
	if requestErrorCode != nil {
		return nil, requestErrorCode
	}

	merchantResult := merchantPersistence.GetByMerchantIdOrNil(request.MerchantId)
	if merchantResult != nil {
		return nil, exception.BadRequestError(string(exception.UNAVAILABLE_MERCHANT_ID))
	}

	var merchant merchant.Merchant
	merchant.Id = 0
	merchant.MerchantId = request.MerchantId
	merchant.Name = request.Name
	merchant.CreatedAt = time.Now().UnixMilli()
	merchant.Code = fmt.Sprintf("MER_%d", merchant.CreatedAt)

	merchantPersistence.Create(&merchant)

	return &merchant, nil
}
