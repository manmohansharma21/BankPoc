package RestProject1

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDB(dbUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{}) //Config should ensure that settings like re-creation of tables etc.
	if err != nil {
		return nil, err
	}

	return db, err
}

func initializeDB() (*gorm.DB, error) {
	protocol := "postgres" //os.Getenv("PROTOCOL") --> URI scheme that identifies the protocol and driver being used for the connection.
	dbUser := "postgres"   //os.Getenv("DB_USER")
	dbPasswd := "postgres" // os.Getenv("DB_PASSWD")
	dbAddr := "localhost"  //os.Getenv("DB_ADDR")
	dbPort := "5430"       //os.Getenv("DB_PORT")
	dbName := "gobank"     //os.Getenv("DB_NAME")

	//db configuration before the server
	// Example:- dbUrl := "postgres://postgres:postgres@localhost:5430/gobank" //avoiding to use 5432 port as there can be some conflict with already running db on default port
	//gobank is the database name

	dbUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s", protocol, dbUser, dbPasswd, dbAddr, dbPort, dbName) // Example:- "root:root@tcp(localhost:3306)/banking"

	db, err := getDB(dbUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}

/*
To connect to a PostgreSQL database using the GORM library and want
 to construct the database URL dynamically from environment variables.
 GORM library and the PostgreSQL driver (gorm.io/driver/postgres) ,
*/
