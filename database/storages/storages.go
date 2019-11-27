package database

import (
	"context"
	"fmt"
	"log"

	models "github.com/whoismath/go-telegram-bot-base/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handler is the structure containing all data types used for store values in database
type Handler struct {
	DB     *mongo.Client
	Domain models.DomainModel
}

// CreateStorages is a function that setup parameters to start storing entries
func CreateStorages(c *mongo.Client) (*Handler, error) {
	return &Handler{
		DB:     c,
		Domain: models.DomainModel{},
	}, nil
}

// SetDomain is a function that receive a string and store it as domain in Handler structure (isn't the best way to do that i know)
func (s *Handler) SetDomain(d string) error {
	domain := models.DomainModel{primitive.NewObjectID(), primitive.NilObjectID, d}
	s.Domain = domain
	return nil
}

// SetSubdomain is a function that receive a string with a subdomain and store it as DOMAIN in Handler structure
func (s *Handler) SetSubdomain(i primitive.ObjectID, d string) error {
	domain := models.DomainModel{primitive.NewObjectID(), i, d}
	s.Domain = domain
	return nil
}

// CreateDomain receives a domain and insert it into databasee returning any error
func (s *Handler) CreateDomain(d string) error {

	collection := s.DB.Database("go-telegram-bot-base-bot").Collection("domains")

	s.SetDomain(d)
	insertResult, err := collection.InsertOne(context.TODO(), s.Domain)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertResult)

	return nil
}

// CreateSubdomain receives a subdomain and store it into database returning any error
func (s *Handler) CreateSubdomain(i primitive.ObjectID, d string) error {

	collection := s.DB.Database("go-telegram-bot-base-bot").Collection("domains")

	s.SetSubdomain(i, d)
	_, err := collection.InsertOne(context.TODO(), s.Domain)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// DeleteDomain receives a domains and delete it from the database
func (s *Handler) DeleteDomain(i primitive.ObjectID) (bool, error) {
	collection := s.DB.Database("go-telegram-bot-base-bot").Collection("domains")

	_, err := collection.DeleteMany(context.TODO(), bson.D{{"_id", i}})
	if err != nil {
		log.Fatal(err)
	}
	_, err = collection.DeleteMany(context.TODO(), bson.D{{"subdomain", i}})
	if err != nil {
		log.Fatal(err)
	}

	return true, nil
}
