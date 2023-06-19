package util

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/IkehAkinyemi/myblog/internal/models"
)

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
type TemplateData struct {
	CurrentYear     int
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
	Article Article
}

type Article struct {
	models.Post
	Content template.HTML
}

var functions = template.FuncMap{
	"friendlyDataFormat": FriendlyDataFormat,
}

func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	// Initializw cache
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*layout.html"))
		if err != nil {
			return nil, err
		}

		// ts, err = ts.ParseGlob(filepath.Join(dir, "*partial.html"))
		// if err != nil {
		// 	return nil, err
		// }

		cache[name] = ts
	}

	return cache, nil
}

func FriendlyDataFormat(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}
