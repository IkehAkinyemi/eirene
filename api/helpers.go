package api

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/IkehAkinyemi/myblog/internal/util"
)

// The serverError helper writes an error message and stack trace to the errorLog
func (s *Server) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	s.logger.Error().Str("stack-trace", trace).Msg("server error occurred")

	if s.configs.Env == "development" {
		http.Error(w, trace, http.StatusInternalServerError)
		return
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends s specfic status code and
// corresponding description to http.ResponseWriter
func (s *Server) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For convienence, we'll implement a notFound error response
func (s *Server) notFound(w http.ResponseWriter) {
	s.clientError(w, http.StatusNotFound)
}

// render retrieves and render the appropriate template set from the cache based on the page name
func (s *Server) render(w http.ResponseWriter, r *http.Request, name string, td *util.TemplateData) {
	ts, ok := s.templateCache[name]
	if !ok {
		s.serverError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}
	buf := new(bytes.Buffer)
	err := ts.Execute(buf, s.addDefaultData(td, r))
	if err != nil {
		s.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}

func (s *Server) addDefaultData(td *util.TemplateData, r *http.Request) *util.TemplateData {
	if td == nil {
		td = &util.TemplateData{}
	}
	td.CurrentYear = time.Now().Year()
	return td
}
