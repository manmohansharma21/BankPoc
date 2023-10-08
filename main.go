package main

func main() {
	//fmt.Println("=============Jai Jai RadhaKrishn========")
	//db before the server
	dbUrl := "postgres://postgres:postgres@localhost:5430/gobank" //avoiding to use 5432 port as there can be some conflict with already running db on default port
	//gobank is the database name
	db, err := getDB(dbUrl)
	if err != nil {
		panic("failed to connect to the database")
	}

	addr := ":3029"
	server1 := newAPIServer(addr, db) //server2:= newAPIServer(":4000")
	server1.Run()

}

/*
Receiver receives the dependencies, that is why its name, to provide dependency injection.
*/
