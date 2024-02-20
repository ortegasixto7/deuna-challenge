package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortegasixto7/deuna-challenge/bank/src/core/customer"
	"github.com/ortegasixto7/deuna-challenge/bank/src/persistence/postgres"
	"github.com/ortegasixto7/deuna-challenge/bank/src/usecases/customer/deposit"
	"github.com/ortegasixto7/deuna-challenge/bank/src/usecases/customer/signup"
)

type CustomerController struct {
	CustomerPersistence *customer.CustomerPersistence
}

func (controller *CustomerController) SignUp(c *fiber.Ctx) error {
	request := new(signup.SignUpRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}
	data, err := signup.Execute(request, new(postgres.PostgresCustomerPersistence))
	if err != nil {
		return err
	}
	c.JSON(data)
	return nil
}

func (controller *CustomerController) Deposit(c *fiber.Ctx) error {
	request := new(deposit.DepositRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}
	data, err := deposit.Execute(request, new(postgres.PostgresCustomerPersistence))
	if err != nil {
		return err
	}
	c.JSON(data)
	return nil
}
