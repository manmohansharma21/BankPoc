package RestProject1

import (
	"fmt"
	"log"
	"os"
)

// sanityCheck: We check for env variables, and if any var is empty(not defined), we do not start the application.
func sanityCheck() {
	if os.Getenv("SERVER_PORT") == "" { //Can add more env variable, and can print the name of missing ones.
		log.Fatal("Environment variable not defined")
	}
}
func Start() { //If we do not expose by capitalizing, it will not be accessible outside app directory.

	fmt.Println("=============Jai Jai Shree RadhaKrishn========")

	port := os.Getenv("SERVER_PORT") //port := ":3078"
	address := ":" + port

	db, err := initializeDB()
	if err != nil {
		panic("failed to connect to the database")
	}

	server1 := newAPIServer(address, db) // Ex, "8080" is the port but ":8080" with colon is the server. server2:= newAPIServer(":4000")
	server1.Run()
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
