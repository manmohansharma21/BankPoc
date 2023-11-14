package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/manmohansharma21/bankpoc/banking-lib/errs"
	"github.com/manmohansharma21/bankpoc/banking-lib/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) { // When read only use of receiver, no need of pointer.
	sqlInsert := "INSERT INTO accounts(customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	// This above query returns lastInsertedId which is customer_id in our case.
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10) // or a.AccountId = string(id)
	return &a, nil
}

/*
* transaction = make an entry in the transaction table + update the balance in the accounts table
 */
func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date)
	values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// commit the transaction when all is good.
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while committing transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Getting the last transaction_id from the transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected databse error")
	}

	// Getting the latest account information from the accounts table
	account, appErr := d.FindById(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// Updating the transactio struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDb) FindById(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account

	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &account, nil
}

// NewAccountRepositoryDb: Helper function to create an instance of AccountRepositoryDb
func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}

/*
Choosing Between Pointer and Non-Pointer Receivers:
Consider whether the method needs to modify the original value or if it only needs to read it.
If you want to modify the original value, use a pointer receiver.
If you want to work with an immutable value or avoid accidental modification, use a non-pointer receiver.
Value Semantics vs. Reference Semantics:
Methods with non-pointer receivers use value semantics, which means they operate on copies of the value. Any modifications are isolated to the method's scope.
Methods with pointer receivers use reference semantics, which means they operate on the original value by reference. Modifications affect the original value directly.
*/
