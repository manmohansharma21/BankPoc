package RestProject1

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// APIServer represents an API server.
// It may be injected as a dependency and may contain information such as a database connection,
// an HTTP client, AWS information, etc.
type APIServer1 struct {
	addr    string   // The server address.
	storage *Storage // Database connection or GORM instance. Storage layer or use directly db   *gorm.DB
	// Additional fields for HTTP client, AWS information, etc.
	client *Client
}

/*
The reason we create the Storage and Client wrappers instead of exposing *gorm.DB and http.Client directly
is to restrict access to the database and HTTP client functionalities to their respective related functions/methods.
This encapsulation also enables effective plugged in as dependency injection, ensuring that these components are used only
within their intended contexts. Similiarly, for the client here. Hence, not exposing everthing to server but to concerned services only.
Storage and Client are indeed acting as wrappers around the *gorm.DB and http.Client instances.
In software development, a "wrapper" is a common term used to describe a component or interface that encapsulates or "wraps" another component to provide a different or more controlled interface. These wrappers can add additional functionality or restrict access to the underlying component in a specific way.

Can hit any endpoint either using 'curl' command on terminal or localhost over the browser or on Postman app.
routes==endpoints==patterns==apis on high level.
manmohansharma@Manmohans-MacBook-Air ~ % curl http://localhost:3055/customers
[{"full_name":"Manmohan","city":"Vrindavan","zipcode":"281121"},{"full_name":"Bhavesh","city":"Vrindavan","zipcode":"281121"}]

*/

// Helper Function
func newAPIServer(addr string, db *gorm.DB) *APIServer1 {
	httpClient := &http.Client{}
	return &APIServer1{
		addr: addr,
		storage: &Storage{
			db: db,
		}, // db connection is passed as dependency injected so that we come to know
		// its existence in advance rather than reported by client from the front-end.
		client: &Client{
			httpClient: httpClient,
		},
	}
}

// Every server must have a Run()
func (s *APIServer1) Run() { // Dependency injection via receiver argument
	fmt.Println("Server1 started on port....", s.addr)
	//migrate(): to create table schemas in DB, to be run once in the beginning
	//s.migrate()

	router := mux.NewRouter()

	//Adding Routes to Endpoints
	router.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		s.handleGreet(w, r)
	})
	router.HandleFunc("/login", s.handleLogin)
	router.HandleFunc("/introspect", s.handleIntrospect)
	router.HandleFunc("/post", s.handleGetPost)          //adding handlers to the routes to endpoints
	router.HandleFunc("/post/{id}", s.handleGetPostById) //URL segmenting, retrieving information URL
	router.HandleFunc("/post/createPayload", s.handleCreatePost)
	//s.interruptDemo()

	fmt.Printf("Server starting on address %s", s.addr)

	//Starting server
	http.ListenAndServe(s.addr, router) //or http.ListenAndServe("localhost:8000", nil); To Start the Server, spinning up the server.
}

// migrate(): to create table schemas in DB, to be run once in the beginning
func (s *APIServer1) migrate() {
	// Migrate the schema
	s.storage.db.AutoMigrate(&Post{})
}
