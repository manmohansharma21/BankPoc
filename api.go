package main

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
}

func newAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		storage: &Storage{
			db: db,
		}, // db connection is passed as dependency injected so that we come to know
		// its existence in advance rather than reported by client from the front-end.
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
