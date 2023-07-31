package api

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/IkehAkinyemi/eirene/models"
	tmplcache "github.com/IkehAkinyemi/eirene/tmplCache"
	"github.com/rs/zerolog/log"
	"github.com/russross/blackfriday/v2"
	"gopkg.in/yaml.v2"
)

func (s *Server) redir2Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
}

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, "home.page.html", nil)
}

func (s *Server) posts(w http.ResponseWriter, r *http.Request) {
	var posts models.Posts

	files, err := os.ReadDir(s.articleDir)
	if err != nil {
		log.Fatal().Err(err).Msg("error occurred")
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			content, err := os.ReadFile(filepath.Join(s.articleDir, file.Name()))
			if err != nil {
				log.Fatal().Err(err).Msg("error occurred")
			}

			var post models.Post
			parts := bytes.SplitN(content, s.sep, 3)
			if len(parts) == 3 {
				err = yaml.Unmarshal(parts[1], &post)
				if err != nil {
					log.Fatal().Err(err).Msg("error occurred")
				}
				post.Content = string(blackfriday.Run(parts[2]))
				posts = append(posts, post)
			}
		}
	}

	sort.Sort(posts)

	td := &tmplcache.TemplateData{
		Articles: posts,
	}

	s.render(w, r, "posts.page.html", td)
}

func (s *Server) post(w http.ResponseWriter, r *http.Request) {
	// Get the "id" parameter from the route
	id := r.URL.Query().Get(":id")
	filename := fmt.Sprintf("%s.md", id)

	article, err := os.ReadFile(filepath.Join(s.articleDir, filename))
	if err != nil {
		log.Fatal().Err(err).Msg("error occurred")
	}

	parts := bytes.SplitN(article, s.sep, 3)
	if len(parts) != 3 {
		log.Fatal().Err(err).Msg("invalid markdown format")
	}

	var post models.Post
	err = yaml.Unmarshal(parts[1], &post)
	if err != nil {
		log.Fatal().Err(err).Msg("error occurred")
	}

	post.Content = string(blackfriday.Run(parts[2]))

	td := &tmplcache.TemplateData{
		Article: tmplcache.Article{
			Post:    post,
			Content: template.HTML(post.Content),
		},
	}
	s.render(w, r, "post.page.html", td)
}
