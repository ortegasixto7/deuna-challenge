package postgres

import (
	"fmt"

	"github.com/ortegasixto7/deuna-challenge/bank/src/core/transaction"
)

type PostgresTransactionPersistence struct{}

func (persistence *PostgresTransactionPersistence) Create(data *transaction.Transaction) {
	result := Database.Create(&data)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Println("Transaction Created!")
	}
}

func (persistence *PostgresTransactionPersistence) Update(data *transaction.Transaction) {
	result := Database.Save(data)
	if result.Error != nil {
		fmt.Println("Error updating Transaction ....")
		fmt.Println(result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Println("Transaction Updated!")
	}
}

func (persistence *PostgresTransactionPersistence) GetByCodeOrNil(code string) *transaction.Transaction {
	var result *transaction.Transaction
	Database.Where(&transaction.Transaction{Code: code}).First(&result)
	if result.Id == 0 {
		return nil
	}
	return result
}
