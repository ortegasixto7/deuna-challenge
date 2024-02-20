package postgres

import (
	"fmt"

	"github.com/ortegasixto7/deuna-challenge/payment-platform/src/core/merchant"
)

type PostgresMerchantPersistence struct{}

func (persistence *PostgresMerchantPersistence) Create(data *merchant.Merchant) {
	result := Database.Create(&data)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Println("Merchant Created!")
	}
}

func (persistence *PostgresMerchantPersistence) Update(data *merchant.Merchant) {
	result := Database.Save(data)
	if result.Error != nil {
		fmt.Println("Error updating Merchant ....")
		fmt.Println(result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Println("Merchant Updated!")
	}
}

func (persistence *PostgresMerchantPersistence) GetByMerchantIdOrNil(merchantId string) *merchant.Merchant {
	var result *merchant.Merchant
	Database.Where(&merchant.Merchant{MerchantId: merchantId}).First(&result)
	if result.Id == 0 {
		return nil
	}
	return result
}

func (persistence *PostgresMerchantPersistence) GetByCodeOrNil(code string) *merchant.Merchant {
	var result *merchant.Merchant
	Database.Where(&merchant.Merchant{Code: code}).First(&result)
	if result.Id == 0 {
		return nil
	}
	return result
}
