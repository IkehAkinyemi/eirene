package api

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/IkehAkinyemi/myblog/internal/models"
	"github.com/IkehAkinyemi/myblog/internal/util"
	"github.com/russross/blackfriday/v2"
)

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	// file, err := os.ReadFile("./articles/trial.md")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// output := fmt.Sprintf(`<!DOCTYPE html>
	// <html lang="en">
	// <head>
	// 	<meta charset="UTF-8">
	// 	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	// 	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	// 	<title>Document</title>
	// </head>
	// <body>
	// 	%s
	// </body>
	// </html>`, string(blackfriday.Run(file)))
	// tmpl := template.New("ah!")
	// tmpl.Parse(string(output))
	// tmpl.Execute(w, "")
	s.render(w, r, "home.page.html", nil)
}

func (s *Server) posts(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, "posts.page.html", nil)
}

func (s *Server) post(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("./articles/trial.md")
	if err != nil {
		log.Fatal(err)
	}
	output := blackfriday.Run(file)
	td := &util.TemplateData{
		Article: util.Article{
			Post: models.Post{
				Author: "Ikeh Chukwuka X",
				Title: "Rust Concurrency Patterns for Parallel Programming",
				CreatedAt: time.Now(),
				Tags: []string{"rust", "concurrency", "channels", "threads"},
			},
			Content: template.HTML(string(output)),
		},
	}
	s.render(w, r, "post.page.html", td)
}
