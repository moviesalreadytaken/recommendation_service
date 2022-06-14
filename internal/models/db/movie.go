package db

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ExternalId  uuid.UUID          `bson:"externalId,omitempty"`
	Name        string             `bson:"name,omitempty"`
	ReleaseDate time.Time          `bson:"releaseDate,omitempty"`
	MinAge      uint32             `bson:"minAge,omitempty"`
	Description string             `bson:"description,omitempty"`
	Rate        float32            `bson:"rate,omitempty"`
	Url         string             `bson:"url,omitempty"`
}
