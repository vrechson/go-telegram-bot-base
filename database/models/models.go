package database

import "go.mongodb.org/mongo-driver/bson/primitive"

// DomainStorage aaa
type DomainModel struct {
	ID          primitive.ObjectID `bson:"_id"`
	IsSubdomain primitive.ObjectID `bson:"subdomain"`
	Value       string             `bson:"value"`
}
