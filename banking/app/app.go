package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

// sanityCheck: We check for env variables, and if any var is empty(not defined), we do not start the application.
func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" { //Can add more env variable, and can print the name of missing ones.
		log.Fatal("Environment variable not defined")
	}
}
func Start() { //If we do not expose by capitalizing, it will not be accessible outside app directory.

	fmt.Println("=============Jai Jai Shree RadhaKrishn========")

	sanityCheck()

	address := os.Getenv("SERVER_ADDRESS") // addr:="localhost"
	port := os.Getenv("SERVER_PORT")       //port := ":3078"

	//port = strings.TrimLeft(port, ":") //If required only "8088" not ":8088" this time, no colon as colon is inserted from the fmt.Sprintf
	server2 := newCustomerAPIServer(address, port)
	server2.Run()
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName) // Example:- "root:root@tcp(localhost:3306)/banking"

	//db or client is the client for mysql. "mysql" is the driver name to be used by database/sql to connect with database. Database configuration.
	dbClient, err := sqlx.Open("mysql", dataSource) // example:-"root:root@tcp(localhost:3306)/banking") //Our data-source settings, user:password@/dbname
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	dbClient.SetConnMaxLifetime(time.Minute * 3) //Connetion time span max allowed
	dbClient.SetMaxOpenConns(10)                 //Max limit of open connections allowed
	dbClient.SetMaxIdleConns(10)                 //Max no of connections always available in the connection pool

	return dbClient
}

/*
* Receiver receives the dependencies, that is why its name, to provide dependency injection we attach receivers rather than pssing value via arguments.
* ORM simplifies the DB interactions by providing drivers which have direct functions using GO structures rather than specifying the SQL queries.
* Layers to Manintain During Development:
1. Application Layer
2. Storage Layer
3. Authentication Layer
4. Service OR Infrastructure Layer
*
*/
