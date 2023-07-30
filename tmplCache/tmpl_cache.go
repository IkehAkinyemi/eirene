package tmplcache

import (
	"html/template"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/IkehAkinyemi/eirene/models"
)

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
type TemplateData struct {
	CurrentYear int
	Article     Article
	Articles    []models.Post
}

type Article struct {
	models.Post
	Content template.HTML
}

var functions = template.FuncMap{
	"friendlyDataFormat": FriendlyDataFormat,
	"slugify":            Slugify,
	"calReadTime":        CalculateReadTime,
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

		cache[name] = ts
	}

	return cache, nil
}

func FriendlyDataFormat(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006")
}

func Slugify(s string) string {
	s = strings.ToLower(s)

	s = strings.Map(func(r rune) rune {
		if r > 127 {
			return -1
		}
		return r
	}, s)

	reg := regexp.MustCompile(`[^a-z0-9]+`)
	s = reg.ReplaceAllString(s, "-")

	s = strings.Trim(s, "-")

	return s
}

func CalculateReadTime(text string) int {
	words := strings.Fields(text)
	wordCount := len(words)

	// Assume an average reading speed of 200 words per minute
	readingSpeed := 200

	readTime := (wordCount + readingSpeed - 1) / readingSpeed // This ensures rounding up

	return readTime
}
