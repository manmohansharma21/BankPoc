package app

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

// Post represents a post entity. Model for the DB.
type Post struct {
	ID        int       `json:"id" gorm:"primaryKey"`         //  //decorating with the JSON tags with which names we need to expose the fields.	The unique identifier for the post.
	Title     string    `json:"title" gorm:"unique;not null"` //  Separate JSON and GORM tags with spaces. The title of the post.
	Content   string    `json:"content"`                      // The content of the post.
	CreatedAt time.Time `json:"created_at"`                   // The timestamp when the post was created.
}

// Model for the JSON payload.
type PostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostRepository interface { //PostRepository to follow Go's convention of using nouns for interface names.
	CreatePost(post *Post) (*Post, error) //camel case
}

type Storage struct {
	db *gorm.DB
}

type KLoginPayload struct {
	clientId     string
	username     string
	password     string
	grantType    string
	clientSecret string
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Client struct {
	httpClient *http.Client
}

type KLoginRes struct {
	AccessToken string `json:"access_token"`
}

type LoginRes struct {
	AccessToken string `json:"access_token"`
}

type KIntrospectPayload struct {
	clientId     string
	clientSecret string
	token        string
}

type IntrospectPayload struct {
	AccessToken string `json:"access_token"`
}

type introspectRes struct {
	// Body interface{} `json:"body"`
	Active bool `json:"active"`
}
