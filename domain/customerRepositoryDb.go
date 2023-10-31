package domain

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // go-sql-driver/mysql driver for Golang to connect to MYSQL
	"github.com/jmoiron/sqlx"
	"github.com/manmohansharma21/bankpoc/errs"
	"github.com/manmohansharma21/bankpoc/logger"
)

type customerRepositoryDb struct {
	dbClient *sqlx.DB //DB Client as a state
}

func (d customerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		//	rows, err = d.dbClient.Query(findAllSql)
		err = d.dbClient.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		//rows, err = d.dbClient.Query(findAllSql, status)
		err = d.dbClient.Select(&customers, findAllSql, status) //Third is the arguments to the select query.
	}

	if err != nil {
		logger.Error("error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexepected database error")
	}

	// err = sqlx.StructScan(rows, customers)
	// if err != nil {
	// 	logger.Error("error while scanning customers" + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexepected database error")
	// }

	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	// 	if err != nil {
	// 		logger.Error("error while scanning customers" + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unexepected database error")
	// 	}

	// 	customers = append(customers, c)
	// }

	return customers, nil
}

func (d customerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?"

	//row := d.dbClient.QueryRow(customerSql, id)

	var c Customer
	err := d.dbClient.Get(&c, customerSql, id)
	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			// return nil, errors.New("customer not found")
			return nil, errs.NewNotFoundError("customer not found")
		}

		log.Println("error while scanning customer" + err.Error())
		// return nil, errors.New("unexpected database error")
		return nil, errs.NewUnexpectedError("unexpected database error")

	}

	return &c, nil
}

// Helper function
func NewCustomerRepositoryDb(dbClient *sqlx.DB) customerRepositoryDb {
	return customerRepositoryDb{dbClient: dbClient}
}

/* Sqlx: Third part library, general purpose extension to database/sql,
helps in removing boiler plate code;
helps to marshal rows into structs.
*/
