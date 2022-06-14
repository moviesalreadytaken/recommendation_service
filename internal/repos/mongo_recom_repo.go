package repos

import (
	models "github.com/moviesalreadytaken/recommendation_service/internal/models/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieRecommendationRepo interface {
	SaveRecommendation(models.MoviesRecommendation) (primitive.ObjectID,error)

	GetAllRecommendations() ([]models.MoviesRecommendation,error)
}
