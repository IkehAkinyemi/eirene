package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/russross/blackfriday/v2"
)

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("./articles/trial.md")
	if err != nil {
		log.Fatal(err)
	}
	output := fmt.Sprintf(`<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		%s
	</body>
	</html>`, string(blackfriday.Run(file)))
	tmpl := template.New("ah!")
	tmpl.Parse(string(output))
	tmpl.Execute(w, "")
	// fmt.Fprintln(w, string(output))
}