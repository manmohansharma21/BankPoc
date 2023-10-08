package main

import (
	"time"

	"gorm.io/gorm"
)

// Post represents a post entity.
type Post struct {
	ID        int       `json:"id" gorm:"primaryKey"`         //  //decorating with the JSON tags with which names we need to expose the fields.	The unique identifier for the post.
	Title     string    `json:"title" gorm:"unique;not null"` //  Separate JSON and GORM tags with spaces. The title of the post.
	Content   string    `json:"content"`                      // The content of the post.
	CreatedAt time.Time `json:"created_at"`                   // The timestamp when the post was created.
}

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
