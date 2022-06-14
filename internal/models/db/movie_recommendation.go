package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MoviesRecommendation struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Movies             []Movie            `bson:"movies,omitempty"`
	User               User               `bson:"user,omitempty"`
	RecommendationDate time.Time          `bson:"recommendationDate,omitempty"`
	Description        string             `bson:"description,omitempty"`
}
