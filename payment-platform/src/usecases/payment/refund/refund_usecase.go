package refund

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/payment"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/exception"
)

func Execute(request *RefundRequest, paymentPersistence payment.PaymentPersistence) (*payment.Payment, error) {
	requestErrorCode := request.Validate()
	if requestErrorCode != nil {
		return nil, requestErrorCode
	}

	payment := paymentPersistence.GetByCodeOrNil(request.PaymentCode)
	if payment == nil {
		return nil, exception.NotFoundError(string(exception.PAYMENT_NOT_FOUND))
	}
	if payment.IsRefunded {
		return nil, exception.BadRequestError(string(exception.PAYMENT_IS_ALREADY_REFUNDED))
	}

	bankUrl := os.Getenv("BANK_HOST")
	var GO_ENV string = os.Getenv("GO_ENV")
	if GO_ENV == "docker" {
		bankUrl = os.Getenv("BANK_HOST_DOCKER")
	}
	postBody, _ := json.Marshal(map[string]interface{}{
		"transaction_code": payment.ExternalCode,
	})
	response, err := http.Post(fmt.Sprintf("%s/transactions/refunds/v1", bankUrl), "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return nil, exception.BadRequestError(string(exception.REFUND_FAILED))
	}
	var responseBody BankResponse
	if decErr := json.NewDecoder(response.Body).Decode(&responseBody); decErr != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}

	if response.StatusCode != http.StatusOK {
		if responseBody.Error == string(exception.PAYMENT_IS_ALREADY_REFUNDED) {
			return nil, exception.BadRequestError(string(exception.PAYMENT_IS_ALREADY_REFUNDED))
		}
		return nil, exception.BadRequestError(string(exception.REFUND_FAILED))
	}
	defer response.Body.Close()

	payment.IsRefunded = true
	paymentPersistence.Update(payment)

	return payment, nil
}
