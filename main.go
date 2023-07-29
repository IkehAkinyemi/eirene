package main

import (
	"os"
	"time"

	"github.com/IkehAkinyemi/eirene/api"
	"github.com/IkehAkinyemi/eirene/configs"
	tmplcache "github.com/IkehAkinyemi/eirene/tmplCache"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	configs, err := configs.ParseConfigs("./")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse configurations")
	}

	if configs.Env == "development" {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC3339,
			},
		).With().Caller().Logger()
	}

	templateCache, err := tmplcache.NewTemplateCache("./ui/html")
	if err != nil {
		log.Fatal().Err(err).Msg("error occurred")
	}

	server := api.NewServer(configs, templateCache, log.Logger)

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("error occur starting server")
	}
}
