package app

import (
	"fmt"
)

func Start() { //If we do not expose by capitalizing, it will not be accessible outside app directory.

	fmt.Println("=============Jai Jai Shree RadhaKrishn========")

	//db before the server
	dbUrl := "postgres://postgres:postgres@localhost:5430/gobank" //avoiding to use 5432 port as there can be some conflict with already running db on default port
	//gobank is the database name

	db, err := getDB(dbUrl)
	if err != nil {
		panic("failed to connect to the database")
	}

	addr := ":3055"
	server1 := newAPIServer(addr, db) //server2:= newAPIServer(":4000")
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
