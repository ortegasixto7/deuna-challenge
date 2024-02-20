package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortegasixto7/deuna-challenge/bank/src/core/merchant"
	"github.com/ortegasixto7/deuna-challenge/bank/src/persistence/postgres"
	"github.com/ortegasixto7/deuna-challenge/bank/src/usecases/merchant/signup"
)

type MerchantController struct {
	MerchantPersistence *merchant.MerchantPersistence
}

func (controller *MerchantController) SignUp(c *fiber.Ctx) error {
	request := new(signup.SignUpRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}
	data, err := signup.Execute(request, new(postgres.PostgresMerchantPersistence))
	if err != nil {
		return err
	}
	c.JSON(data)
	return nil
}
