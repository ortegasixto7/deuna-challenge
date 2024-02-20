package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/customer"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/merchant"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/payment"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/exception"
)

func Execute(
	request *PaymentRequest,
	customerPersistence customer.CustomerPersistence,
	merchantPersistence merchant.MerchantPersistence,
	paymentPersistence payment.PaymentPersistence) (*payment.Payment, error) {

	fmt.Println()
	requestErrorCode := request.Validate()
	if requestErrorCode != nil {
		return nil, requestErrorCode
	}

	customer := customerPersistence.GetByCodeOrNil(request.CustomerCode)
	if customer == nil {
		return nil, exception.NotFoundError(string(exception.CUSTOMER_NOT_FOUND))
	}

	merchant := merchantPersistence.GetByCodeOrNil(request.MerchantCode)
	if merchant == nil {
		return nil, exception.NotFoundError(string(exception.MERCHANT_NOT_FOUND))
	}

	bankUrl := os.Getenv("BANK_HOST")
	var GO_ENV string = os.Getenv("GO_ENV")
	if GO_ENV == "docker" {
		bankUrl = os.Getenv("BANK_HOST_DOCKER")
	}
	postBody, _ := json.Marshal(map[string]interface{}{
		"customer_id": customer.NationalId,
		"merchant_id": merchant.MerchantId,
		"amount":      request.Amount,
		"card_token":  request.CardToken,
	})
	response, err := http.Post(fmt.Sprintf("%s/transactions/transfers/v1", bankUrl), "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return nil, exception.BadRequestError(string(exception.PAYMENT_FAILED))
	}
	var responseBody BankResponse
	if decErr := json.NewDecoder(response.Body).Decode(&responseBody); decErr != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}
	if response.StatusCode != http.StatusOK {
		if responseBody.Error == string(exception.INSUFFICIENT_FUNDS) {
			return nil, exception.BadRequestError(string(exception.INSUFFICIENT_FUNDS))
		}
		if responseBody.Error == "INVALID_CARD_TOKEN" {
			return nil, exception.BadRequestError(string(exception.INVALID_CARD_TOKEN))
		}
		if responseBody.Error == "INVALID_TRANSACTION" {
			return nil, exception.BadRequestError(string(exception.PAYMENT_DECLINED))
		}
		return nil, exception.BadRequestError(string(exception.PAYMENT_FAILED))
	}
	defer response.Body.Close()

	var payment payment.Payment
	payment.Id = 0
	payment.Code = fmt.Sprintf("PAY_%d", time.Now().UnixMilli())
	payment.CustomerId = customer.NationalId
	payment.CustomerName = customer.Name
	payment.MerchantId = merchant.MerchantId
	payment.MerchantName = merchant.Name
	payment.Amount = request.Amount
	payment.IsRefunded = false
	payment.CreatedAt = time.Now().UnixMilli()
	payment.ExternalCode = responseBody.Code
	paymentPersistence.Create(&payment)

	return &payment, nil
}
