package db

import (
	"github.com/IkehAkinyemi/myblog/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient struct {
	client *mongo.Client
}

func NewMongoClient(c *mongo.Client) models.Store {
	return &MongoClient{
		client: c,
	}
}
