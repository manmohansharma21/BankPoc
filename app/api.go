package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// APIServer represents an API server.
// It may be injected as a dependency and may contain information such as a database connection,
// an HTTP client, AWS information, etc.
type APIServer struct {
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
*/

func newAPIServer(addr string, db *gorm.DB) *APIServer {
	httpClient := &http.Client{}
	return &APIServer{
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
func (s *APIServer) Run() { // Dependency injection via receiver argument
	//migrate(): to create table schemas in DB, to be run once in the beginning
	s.migrate()

	router := mux.NewRouter()

	//Adding Routes to Endpoints
	router.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		s.handleGreet(w, r)
	})
	router.HandleFunc("/login", s.handleLogin)
	router.HandleFunc("/introspect", s.handleIntrospect)
	router.HandleFunc("/post", s.handleGetPost) //adding handlers to the routes to endpoints
	router.HandleFunc("/post/{id}", s.handleGetPostById)
	router.HandleFunc("/post/createPayload", s.handleCreatePost)
	//s.interruptDemo()

	fmt.Printf("Server starting on address %s", s.addr)
	http.ListenAndServe(s.addr, router) //To Start the Server, spinning up the server
}

// migrate(): to create table schemas in DB, to be run once in the beginning
func (s *APIServer) migrate() {
	// Migrate the schema
	s.storage.db.AutoMigrate(&Post{})
}
