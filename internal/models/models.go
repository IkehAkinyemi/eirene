package models

import "time"

// Post models an article.
type Post struct {
	Author string `json:"author" bson:"author"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Tags []string `json:"tags" bson:"tags"`
	Content string `json:"content" bson:"content"`
	Title string `json:"title" bson:"title"`
	Synopsis string `json:"content_synopsis" bson:"content_synopsis"`
	Comments []Comment `json:"comments" bson:"comments"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// Comment models a post's comment.
type Comment struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Content string `json:"content" bson:"content"`
}