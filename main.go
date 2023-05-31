package main

import (
	"log"
	"net/http"

	"github.com/IkehAkinyemi/myblog/cmd/api"
)

func main() {
	// file, err := os.ReadFile("./articles/trial.md")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// output := blackfriday.Run(file)
	s := api.NewServer()
	if err := http.ListenAndServe(":8600", s.Start()); err != nil {
		log.Fatal("couldn't start server")
	}
}