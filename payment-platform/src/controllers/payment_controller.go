package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/customer"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/merchant"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/payment"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/persistence/postgres"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/usecases/payment/getbycode"
	pay "github.com/ortegasixto7/deuna-challenge/payment-platform/src/usecases/payment/payment"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/usecases/payment/refund"
)

type PaymentController struct {
	CustomerPersistence *customer.CustomerPersistence
	MerchantPersistence *merchant.MerchantPersistence
	PaymentPersistence  *payment.PaymentPersistence
}

func (controller *PaymentController) Payment(c *fiber.Ctx) error {
	request := new(pay.PaymentRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}
	data, err := pay.Execute(request, new(postgres.PostgresCustomerPersistence), new(postgres.PostgresMerchantPersistence), new(postgres.PostgresPaymentPersistence))
	if err != nil {
		return err
	}
	c.JSON(data)
	return nil
}

func (controller *PaymentController) Refund(c *fiber.Ctx) error {
	request := new(refund.RefundRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}
	data, err := refund.Execute(request, new(postgres.PostgresPaymentPersistence))
	if err != nil {
		return err
	}
	c.JSON(data)
	return nil
}

func (controller *PaymentController) GetByCode(c *fiber.Ctx) error {
	requestCode := c.Params("code")
	data, err := getbycode.Execute(requestCode, new(postgres.PostgresPaymentPersistence))
	if err != nil {
		return err
	}
	c.JSON(data)
	return nil
}
