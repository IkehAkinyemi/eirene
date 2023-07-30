package models

import "time"

// Post models an article.
type Post struct {
	Author    string    `yaml:"Author"`
	CreatedAt time.Time `yaml:"CreatedAt"`
	Content   string    `json:"content" bson:"content"`
	Title     string    `yaml:"Title"`
	Synopsis  string    `yaml:"Synopsis"`
	ID        int       `yaml:"ArticleID"`
}
