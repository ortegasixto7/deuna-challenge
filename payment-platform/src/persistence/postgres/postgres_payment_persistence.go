package postgres

import (
	"fmt"

	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/payment"
)

type PostgresPaymentPersistence struct{}

func (persistence *PostgresPaymentPersistence) Create(data *payment.Payment) {
	result := Database.Create(&data)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Println("Payment Created!")
	}
}

func (persistence *PostgresPaymentPersistence) Update(data *payment.Payment) {
	result := Database.Save(data)
	if result.Error != nil {
		fmt.Println("Error updating Payment ....")
		fmt.Println(result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Println("Payment Updated!")
	}
}

func (persistence *PostgresPaymentPersistence) GetByCodeOrNil(code string) *payment.Payment {
	var result *payment.Payment
	Database.Where(&payment.Payment{Code: code}).First(&result)
	if result.Id == 0 {
		return nil
	}
	return result
}
