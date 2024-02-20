package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortegasixto7/deuna-challenge/bank/src/core/customer"
	"github.com/ortegasixto7/deuna-challenge/bank/src/core/merchant"
	"github.com/ortegasixto7/deuna-challenge/bank/src/core/transaction"
	"github.com/ortegasixto7/deuna-challenge/bank/src/persistence/postgres"
	"github.com/ortegasixto7/deuna-challenge/bank/src/usecases/transaction/refund"
	"github.com/ortegasixto7/deuna-challenge/bank/src/usecases/transaction/transfer"
)

type TransactionController struct {
	CustomerPersistence    *customer.CustomerPersistence
	MerchantPersistence    *merchant.MerchantPersistence
	TransactionPersistence *transaction.TransactionPersistence
}

func (controller *TransactionController) Transfer(c *fiber.Ctx) error {
	request := new(transfer.TransferRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}
	data, err := transfer.Execute(request, new(postgres.PostgresCustomerPersistence), new(postgres.PostgresMerchantPersistence), new(postgres.PostgresTransactionPersistence))
	if err != nil {
		return err
	}
	c.JSON(data)
	return nil
}

func (controller *TransactionController) Refund(c *fiber.Ctx) error {
	request := new(refund.RefundRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}
	data, err := refund.Execute(request, new(postgres.PostgresCustomerPersistence), new(postgres.PostgresMerchantPersistence), new(postgres.PostgresTransactionPersistence))
	if err != nil {
		return err
	}
	c.JSON(data)
	return nil
}
