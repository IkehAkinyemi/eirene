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

type Posts []Post

func (p Posts) Len() int {
	return len(p)
}

func (p Posts) Less(i, j int) bool {
	return p[i].CreatedAt.After(p[j].CreatedAt) // Sort posts by latest date
}

func (p Posts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
