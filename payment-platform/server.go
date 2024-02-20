package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/controllers"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/exception"
	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/persistence/postgres"
)

func main() {
	var PORT string = ":8080"
	var GO_ENV string = os.Getenv("GO_ENV")
	if GO_ENV == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		PORT = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

	postgres.Init()
	var customerController controllers.CustomerController
	var merchantController controllers.MerchantController
	var paymentController controllers.PaymentController
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			var data = make(map[string]string)

			var badRequest *exception.BadRequestException
			if errors.As(err, &badRequest) {
				data["error"] = err.Error()
				return ctx.Status(fiber.StatusBadRequest).JSON(data)
			}

			var notFound *exception.NotFoundException
			if errors.As(err, &notFound) {
				data["error"] = err.Error()
				return ctx.Status(fiber.StatusNotFound).JSON(data)
			}

			data["error"] = "INTERNAL_ERROR"
			fmt.Println(err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(data)

		},
	})

	app.Get("/", func(c *fiber.Ctx) error {
		var data = make(map[string]string)
		data["message"] = "Payment Platform is running!!"
		return c.JSON(data)
	})

	app.Post("/customers/sign-up/v1", customerController.SignUp)
	app.Post("/merchants/sign-up/v1", merchantController.SignUp)
	app.Post("/payments/payments/v1", paymentController.Payment)
	app.Post("/payments/refunds/v1", paymentController.Refund)
	app.Get("/payments/:code/v1", paymentController.GetByCode)

	fmt.Printf("App running on port: %s", PORT)
	app.Listen(PORT)

}
