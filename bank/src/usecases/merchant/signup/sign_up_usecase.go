package signup

import (
	"time"

	"github.com/ortegasixto7/deuna-challenge/bank/src/core/merchant"
	"github.com/ortegasixto7/deuna-challenge/bank/src/exception"
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
	merchant.Balance = 0
	merchant.CreatedAt = time.Now().UnixMilli()

	merchantPersistence.Create(&merchant)

	return &merchant, nil
}
