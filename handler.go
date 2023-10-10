package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// To expose a struct field for JSON encoding and make it accessible to clients (e.g., Postman), capitalize the field name.
// JSON tags are specified in lowercase and define the JSON keys used in the response.
type GreetResponse struct {
	Greeting string `json:"greeting"`
}

// handleGreet handles HTTP GET requests for the /greet endpoint.
// It responds with a greeting message in JSON format.
// It expects an injected APIServer as a dependency.
func (s *APIServer) handleGreet(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is not GET.
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest) // Respond with a 400 Bad Request status code, i.e., w.WriteHeader(400) 	w.Write([]byte("Method not supported"))
		return
	}

	// Write a greeting message to the response.
	w.Write([]byte("=============Jai Jai SHREERADHAKRISHN============\n"))

	// Set the Content-Type header to application/json.
	// This header informs the client that the response body contains JSON data.
	w.Header().Add("Content-Type", "application/json") // set the Content-Type header to indicate JSON content to be returned.

	// Respond with a 200 OK status code.
	w.WriteHeader(http.StatusOK)

	// Create a GreetResponse struct to hold the greeting message.
	res := &GreetResponse{
		Greeting: "==============Jai Jai ShreeRadhaKrishn============",
	}

	// Encode the GreetResponse struct as JSON and write it to the response body to be sent to the client.
	json.NewEncoder(w).Encode(res) // serialize a Go struct to JSON and write it to the response without need of an intermediate byte slice.
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400) //Response code for bad request
		w.Write([]byte("Method not supported"))
		return
	}

	payload := new(LoginPayload) // Reading from the body ( on Postman for testing)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Payload"))
		return
	}

	kPayload := &KLoginPayload{
		clientId:     "bankpoc-auth",
		username:     payload.Username,
		password:     payload.Password,
		grantType:    "password",
		clientSecret: "gYCWXhHVgWH4ImXfhxmL8C0oAa3mQZUA",
	}

	kres, err := s.client.login(kPayload)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	// res := &LoginRes{
	// 	AccessToken: kres.AccessToken,
	// }

	w.Header().Add("content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(kres)

}

func (s *APIServer) handleIntrospect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400) //Response code for bad request
		w.Write([]byte("Method not supported"))
		return
	}

	payload := new(IntrospectPayload)
	err := json.NewDecoder(r.Body).Decode(payload) // Reading from the body ( on Postman for testing)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Payload"))
		return
	}

	iPayload := &KIntrospectPayload{
		clientId:     "bankpoc-auth",
		clientSecret: "gYCWXhHVgWH4ImXfhxmL8C0oAa3mQZUA",
		token:        payload.AccessToken,
	}

	ires, err := s.client.introspect(iPayload)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(ires)
}

func (s *APIServer) handleGetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400) //Response code for bad request
		w.Write([]byte("Method not supported"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)

	res := &Post{
		ID:        1,
		Title:     "Golang-Bank-POC",
		Content:   "awesome",
		CreatedAt: time.Now(),
	}
	json.NewEncoder(w).Encode(res)
}

func (s *APIServer) handleGetPostById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400) //Response code for bad request
		w.Write([]byte("Method not supported"))
		return
	}

	// Extract the 'id' parameter from the URL using mux.Vars(r)["id"]
	// and convert it to an integer using strconv.Atoi. Any potential error
	// is stored in the 'err' variable.
	id, err := strconv.Atoi(mux.Vars(r)["id"]) //Extract the id from the Params
	if err != nil {                            //handler handling the errors
		w.WriteHeader(400)
		w.Write([]byte("\nValid integer id is required\n"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)

	res, err := s.storage.getPost(id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("\nid does not exist\n"))
		return
	}
	json.NewEncoder(w).Encode(res)
}

// handleCreatePost handles HTTP POST requests for creating a new post.
func (s *APIServer) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is not POST.
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // Respond with a 400 Bad Request status code, i.e., w.WriteHeader(400), i.e., method not allowed
		w.Write([]byte("Method not supported"))
		return
	}

	// Create a new instance of PostPayload to hold the request data.
	payload := new(PostPayload) //=&PostPayload; Create a new PostPayload instance

	// Decode the JSON payload from the request body into the payload struct.
	err := json.NewDecoder(r.Body).Decode(payload) //Need to pass address to PostPayload type
	if err != nil {
		// Handle JSON decoding error.
		w.WriteHeader(http.StatusBadRequest) // w.WriteHeader(400) Respond with a 400 Bad Request status code.
		w.Write([]byte("Invalid Payload"))
		http.Error(w, err.Error(), http.StatusBadRequest) // Log and handle JSON unmarshal error.
		return
	}

	// Log the received payload data for debugging purposes.
	log.Printf("\nReceived JSON payload: Title = %s, Content = %s", payload.Title, payload.Content)
	log.Printf("\nReceived JSON payload: %+v", payload) // log.Print() does not handle structs directly. Instead, it prints the type of the struct and the address in memory, which might not be very informative.
	//	fmt.Print(payload)                                  //Works directly.

	w.Write([]byte("Post creation in progress from the payload provided")) //This gets written on the screen but the JSON gets displayed on the console.

	// Perform processing to create a new post based on the payload.
	// This includes creating a Post struct, setting its fields,
	// and persisting it to the database.

	post := Post{
		Title:     payload.Title,
		Content:   payload.Content,
		CreatedAt: time.Now(),
	}

	// Persist the post to the database using the storage layer.
	err = s.storage.persistPost(&post)
	if err != nil {
		// Handle the error, log it, and respond with an error status code.
		if IsUniqueConstraintViolationError(err) {
			// Handle unique key constraint violation error.
			w.WriteHeader(http.StatusConflict) // 409 Conflict status code.
			w.Write([]byte("\nUnique key constraint violation\n"))
			w.Write([]byte(string(http.StatusConflict)))
		} else {
			// Handle other database errors.
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError) // 500 Internal Server Error. //w.WriteHeader(500); Respond with a 500 Bad Request status code.
			w.Write([]byte("\nSomething went wrong while creating your post structure\n"))
			w.Write([]byte(err.Error()))
			w.Write([]byte(string(http.StatusInternalServerError)))
		}
		return
	}

	// Respond with a success message and a 201 Created status code.
	w.WriteHeader(http.StatusCreated)

	// Respond with a success message.
	w.Write([]byte("\nPost created"))

}

func IsUniqueConstraintViolationError(err error) bool {
	// Check if the error message or type indicates a unique constraint violation.
	return strings.Contains(err.Error(), "unique constraint violation") ||
		strings.Contains(err.Error(), "duplicate key")
}

/*
These lines of code are used to encode and decode JSON data in a Go HTTP server. Let's break down each line and explain their purpose:

1. `json.NewEncoder(w).Encode(res)`

   - `json.NewEncoder(w)`: This part creates a new JSON encoder that will write JSON data to the given `w` (an `http.ResponseWriter`) when the `Encode` method is called.
   - `Encode(res)`: This part encodes the Go struct `res` into a JSON representation and writes it to the `http.ResponseWriter` `w`. It essentially serializes the `res` struct into a JSON format and sends it as the HTTP response body.

   The combination of these two lines is commonly used to send JSON responses in an HTTP handler. It takes a Go struct (`res` in this case), converts it into JSON, and writes it to the HTTP response.

2. `err := json.NewDecoder(r.Body).Decode(payload)`

   - `json.NewDecoder(r.Body)`: This part creates a new JSON decoder that reads JSON data from the `r.Body` (an `io.Reader`). The `r` represents an HTTP request, and `r.Body` is the request's body, which typically contains the incoming JSON data.
   - `.Decode(payload)`: This part decodes the JSON data from `r.Body` and attempts to populate the Go struct `payload` with the parsed data. It essentially deserializes the incoming JSON into the `payload` struct.

   This line is used to process incoming JSON data from an HTTP request. It reads the JSON data from the request's body, parses it, and attempts to map it to the `payload` struct. If successful, the `payload` struct will contain the data sent in the request.

Together, these lines allow you to work with JSON data in an HTTP server. The first line encodes a Go struct and sends it as a JSON response, while the second line decodes incoming JSON data from an HTTP request and populates a Go struct for further processing. These operations are fundamental for building JSON-based APIs in Go.
*/

/*
Hitting on POSTMAN: POST: localhost:3018/post/createPayload
Body=raw
type=json
content-type=application/json
Example payload: {
    "title": "FirstBook",
    "content": "Go"
}
Body gets printed; and console prints the logs we provisioned.

*/

/*
 Using fmt.Print(payload) works because the fmt package is more permissive when it comes to printing values, including structs, pointers, and other types. When you use fmt.Print(payload), it uses Go's default formatting for the value, which often includes field names and their values for structs.

In contrast, log.Print is typically used for logging, and it might not provide the same level of detail when printing complex types like structs. However, as mentioned earlier, you can use log.Printf with the %+v verb to achieve a similar result that includes field names and values.

So, both fmt.Print(payload) and log.Printf("%+v", payload) are valid ways to print the payload struct, with the latter providing more detailed output suitable for debugging and logging.
*/

/*
You can directly write the JSON response to the ResponseWriter without explicitly marshaling it to JSON using json.Marshal. In fact, the code you provided is a valid and efficient way to send JSON responses in a Go web application.

The use of json.NewEncoder(w).Encode(res) is a concise way to serialize a Go struct to JSON and write it to the response. It takes advantage of the json.NewEncoder to efficiently encode and stream the JSON data directly to the response writer without the need for an intermediate byte slice.

Here's why this approach is often preferred:

Efficiency: json.NewEncoder can efficiently encode large JSON responses without the need to store the entire JSON in memory as a byte slice. It serializes the data in chunks, making it memory-efficient.
Streamlined Code: The combination of json.NewEncoder and Encode is concise and idiomatic, making the code more readable and maintainable.
Concurrency: This approach supports streaming and can be beneficial when dealing with long or continuously generated JSON responses in a concurrent environment.
Overall, while you can choose to marshal the struct to JSON and then write it to the response, the approach you provided using json.NewEncoder and Encode is a common and efficient pattern for handling JSON responses in Go web applications. It's a good practice to use it when applicable, especially for large or streaming responses.
*/

/*
Serialization is the process of converting structured data, such as objects or data structures, into a format that can be easily stored, transmitted, or reconstructed later. This serialized format is typically a linear sequence of bytes, often in a specific encoding like JSON, XML, or binary formats.
Common serialization formats include JSON, XML, Protocol Buffers, and more. Each format has its advantages and is suited for different use cases.

In the context of Go and many other programming languages, serialization often involves converting data structures (e.g., structs, objects) into a format that can be stored, transmitted, or saved to disk, and then converting it back into the original data structure when needed. Libraries and tools for serialization, such as the encoding/json package in Go, make this process straightforward.
*/
