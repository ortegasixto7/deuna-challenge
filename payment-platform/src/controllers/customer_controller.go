package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/customer"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/persistence/postgres"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/usecases/customer/signup"
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
