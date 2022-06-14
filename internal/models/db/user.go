package db

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ExternalId uuid.UUID          `bson:"externalId,omitempty"`
	Username   string             `bson:"username,omitempty"`
	Name       string             `bson:"name,omitempty"`
	Surname    string             `bson:"surname,omitempty"`
	Birthdate  time.Time          `bson:"birthdate,omitempty"`
}
