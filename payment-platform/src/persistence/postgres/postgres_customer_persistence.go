package postgres

import (
	"fmt"

	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/customer"
)

type PostgresCustomerPersistence struct{}

func (persistence *PostgresCustomerPersistence) Create(data *customer.Customer) {
	result := Database.Create(&data)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Println("Customer Created!")
	}
}

func (persistence *PostgresCustomerPersistence) Update(data *customer.Customer) {
	result := Database.Save(data)
	if result.Error != nil {
		fmt.Println("Error updating Customer ....")
		fmt.Println(result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Println("Customer Updated!")
	}
}

func (persistence *PostgresCustomerPersistence) GetByNationalIdOrNil(nationalId string) *customer.Customer {
	var result *customer.Customer
	Database.Where(&customer.Customer{NationalId: nationalId}).First(&result)
	if result.Id == 0 {
		return nil
	}
	return result
}

func (persistence *PostgresCustomerPersistence) GetByCodeOrNil(code string) *customer.Customer {
	var result *customer.Customer
	Database.Where(&customer.Customer{Code: code}).First(&result)
	if result.Id == 0 {
		return nil
	}
	return result
}
