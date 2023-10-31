package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/manmohansharma21/bankpoc/errs"
	"github.com/manmohansharma21/bankpoc/logger"
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
