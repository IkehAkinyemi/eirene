package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/russross/blackfriday/v2"
	"gopkg.in/yaml.v2"
)

type Post struct {
	Author    string    `json:"author" bson:"author" yaml:"Author"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" yaml:"CreatedAt"`
	Content   string    `json:"content" bson:"content"`
	Title     string    `json:"title" bson:"title" yaml:"Title"`
	Synopsis  string    `json:"content_synopsis" bson:"content_synopsis" yaml:"Synopsis"`
}

func main() {

	file, err := os.ReadFile("./articles/trial.md")
	if err != nil {
		log.Fatal(err)
	}

	// Split the YAML front matter from the content
	parts := bytes.SplitN(file, []byte("---"), 3)
	if len(parts) != 3 {
		log.Fatal("Invalid Markdown format")
	}

	// Parse the YAML front matter
	var post Post
	err = yaml.Unmarshal(parts[1], &post)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	// Convert Markdown content to HTML
	post.Content = string(blackfriday.Run(parts[2]))

	fmt.Printf("%+v\n", post)
}
