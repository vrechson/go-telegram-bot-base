package database

import (
	"context"
	"log"

	models "github.com/whoismath/go-telegram-bot-base/database/models"
	"go.mongodb.org/mongo-driver/bson"
	primitives "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Handler is the structure containing all data types used for read values from database
type Handler struct {
	DB     *mongo.Client
	Domain models.DomainModel
}

// CreateRepository is a function that setup parameters to start reading entries
func CreateRepository(c *mongo.Client) (*Handler, error) {
	return &Handler{
		DB:     c,
		Domain: models.DomainModel{},
	}, nil
}

// CheckDomain is a function that checks if a domain is in the database
func (s *Handler) CheckDomain(d string) bool {
	collection := s.DB.Database("go-telegram-bot-base-bot").Collection("domains")
	filter := bson.D{{"value", d}}

	var data models.DomainModel
	err := collection.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		return false
	}

	return true
}

// GetDomainID is a function that returns the Id entry of a domain in database
func (s *Handler) GetDomainID(d string) (primitives.ObjectID, error) {
	collection := s.DB.Database("go-telegram-bot-base-bot").Collection("domains")
	filter := bson.D{{"value", d}}

	var data models.DomainModel
	err := collection.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		return primitives.NilObjectID, err
	}

	return data.ID, nil
}

// ListDomains is a function that return a list of domains stored in database
func (s *Handler) ListDomains() ([]string, error) {
	collection := s.DB.Database("go-telegram-bot-base-bot").Collection("domains")
	findOptions := options.Find()
	filter := bson.D{{"subdomain", primitives.NilObjectID}}
	var docs []*models.DomainModel
	var domains []string

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var data models.DomainModel
		if err := cur.Decode(&data); err != nil {
			log.Fatal(err)
		}
		docs = append(docs, &data)
		domains = append(domains, data.Value)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// for a := range domains {
	// 	fmt.Println("\n\n" + domains[a] + "\n\n")
	// }

	cur.Close(context.TODO())
	return domains, nil

}

// ListSubdomains is a function that return a list of subdomains from a specified domain
func (s *Handler) ListSubdomains(i primitives.ObjectID) ([]string, error) {
	collection := s.DB.Database("go-telegram-bot-base-bot").Collection("domains")
	findOptions := options.Find()
	filter := bson.D{{"subdomain", i}}
	var docs []*models.DomainModel
	var subdomains []string

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var data models.DomainModel
		if err := cur.Decode(&data); err != nil {
			log.Fatal(err)
		}
		docs = append(docs, &data)
		subdomains = append(subdomains, data.Value)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// for a := range domains {
	// 	fmt.Println("\n\n" + domains[a] + "\n\n")
	// }

	cur.Close(context.TODO())
	return subdomains, nil

}
