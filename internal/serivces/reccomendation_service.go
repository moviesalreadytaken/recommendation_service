package serivces

import (
	"errors"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/moviesalreadytaken/recommendation_service/internal/models"
	"github.com/moviesalreadytaken/recommendation_service/internal/models/db"
	"github.com/moviesalreadytaken/recommendation_service/internal/repos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieReccomendationService struct {
	repo         repos.MovieRecommendationRepo
	usersClient  *UsersServiceClient
	moviesClient *MoviesServiceClient
}

func NewMovieRecommendationService(
	repo repos.MovieRecommendationRepo,
	usersClient *UsersServiceClient,
	moviesClient *MoviesServiceClient) *MovieReccomendationService {
	return &MovieReccomendationService{
		repo:         repo,
		usersClient:  usersClient,
		moviesClient: moviesClient,
	}
}

func (s *MovieReccomendationService) UserExists(userId uuid.UUID) (bool, error) {
	_, err := s.usersClient.GetUserById(userId)
	if errors.Is(ErrUserNotFound, err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *MovieReccomendationService) GenRecommndation(userId uuid.UUID) ([]db.Movie, error) {
	user, movies, err := s.getUserAndMovies(userId)
	if err != nil {
		return nil, err
	}
	acceptedMovs := make([]db.Movie, 1)
	for i := 0; i < len(movies); i++ {
		minAge := movies[i].MinAge
		age := time.Since(user.Birthdate)
		if minAge <= uint32((age.Seconds() / 31207680)) {
			acceptedMovs = append(acceptedMovs, movies[i])
		}
	}
	sort.Sort(ByRelaseDate(acceptedMovs))
	acceptedMovs = acceptedMovs[1:]
	sort.Sort(ByRate(acceptedMovs))
	acceptedMovs = acceptedMovs[1:]
	recom := db.MoviesRecommendation{
		ID:                 primitive.NilObjectID,
		Movies:             acceptedMovs,
		User:               *user,
		RecommendationDate: time.Now(),
		Description:        "some recommendation description",
	}
	_, err = s.repo.SaveRecommendation(recom)
	if err != nil {
		return nil, err
	}
	return acceptedMovs, nil
}

func (s *MovieReccomendationService) GetAllRecommendations() ([]db.MoviesRecommendation, error) {
	return s.repo.GetAllRecommendations()
}

func (s *MovieReccomendationService) getUserAndMovies(userId uuid.UUID) (*db.User, []db.Movie, error) {
	user, err := s.usersClient.GetUserById(userId)
	if err != nil {
		return nil, nil, err
	}
	dbUser, err := models.RestUserToDb(*user)
	if err != nil {
		return nil, nil, err
	}
	movs, err := s.moviesClient.GetAllMovies()
	if err != nil {
		return nil, nil, err
	}
	dbMovs, err := models.RestMoviesToDb(movs)
	if err != nil {
		return nil, nil, err
	}
	return dbUser, dbMovs, nil
}
