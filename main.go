package main

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/IkehAkinyemi/myblog/api"
	"github.com/IkehAkinyemi/myblog/internal/db"
	"github.com/IkehAkinyemi/myblog/internal/models"
	"github.com/IkehAkinyemi/myblog/internal/util"
	"github.com/rs/zerolog"
)

func main() {
	// file, err := os.ReadFile("./articles/trial.md")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// output := blackfriday.Run(file)

	configs, err := util.ParseConfigs("./")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse configurations")
	}

	if configs.Env == "development" {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: time.RFC3339,
			},
		).With().Caller().Logger()
	}

	


	conn, err := connectDB(configs.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	log.Info().Msg("database connection established")
	defer closeDB(conn)
	store := db.NewMongoClient(conn)

	runServer(configs, store)
}

func runServer(configs util.Configs, store models.Store) {
	server := api.NewServer(configs, store)

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("error occur starting server")
	}
}

// connectDB establishes connection to MongoDB
func connectDB(connURI string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(connURI).
		SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, err
}

// closeDB close database connection.
func closeDB(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client.Disconnect(ctx)
}