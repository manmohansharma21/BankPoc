package main

import "github.com/manmohansharma21/bankpoc/app"

func main() {
	app.Start()
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
