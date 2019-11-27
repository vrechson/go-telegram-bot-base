package factory

import (
	"context"
	"fmt"
	"log"

	"github.com/whoismath/go-telegram-bot-base/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseFactory daawre s
func DatabaseFactory(c *config.Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(c.Db)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client, nil
}
