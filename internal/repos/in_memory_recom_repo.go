package repos

import (
	models "github.com/moviesalreadytaken/recommendation_service/internal/models/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InMemoryRecommendationRepo struct {
	recommendations []models.MoviesRecommendation
}

func NewInMemoryRecommendationRepo() *InMemoryRecommendationRepo {
	return &InMemoryRecommendationRepo{
		recommendations: make([]models.MoviesRecommendation, 0),
	}
}

func (r *InMemoryRecommendationRepo) SaveRecommendation(
	recom models.MoviesRecommendation) (primitive.ObjectID, error) {
	id := primitive.NewObjectID()
	recom.ID = id
	r.recommendations = append(r.recommendations, recom)
	return id, nil
}

func (r *InMemoryRecommendationRepo) GetAllRecommendations() ([]models.MoviesRecommendation, error) {
	return r.recommendations, nil
}
