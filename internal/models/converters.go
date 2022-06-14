package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/moviesalreadytaken/recommendation_service/internal/models/db"
	"github.com/moviesalreadytaken/recommendation_service/internal/models/rest"
)

func DbRecommendationToRest(rec db.MoviesRecommendation) rest.MovieRecommendation {
	return rest.MovieRecommendation{
		ID:                 rec.ID.Hex(),
		Movies:             DbMoviesToRest(rec.Movies),
		User:               DbUserToRest(rec.User),
		RecommendationDate: rec.RecommendationDate,
		Description:        rec.Description,
	}
}

func DbRecommendationsToRest(recs []db.MoviesRecommendation) []rest.MovieRecommendation {
	restRecs := make([]rest.MovieRecommendation, len(recs))
	for i := 0; i < len(recs); i++ {
		restRecs[i] = DbRecommendationToRest(recs[i])
	}
	return restRecs
}

func DbMovieToRest(m db.Movie) rest.Movie {
	return rest.Movie{
		ID:          m.ExternalId.String(),
		Name:        m.Name,
		ReleaseDate: rest.DateOnly(m.ReleaseDate),
		MinAge:      m.MinAge,
		Description: m.Description,
		Rate:        m.Rate,
		Url:         m.Url,
	}
}

func DbMoviesToRest(mvs []db.Movie) []rest.Movie {
	movies := make([]rest.Movie, len(mvs))
	for i := 0; i < len(mvs); i++ {
		movies[i] = DbMovieToRest(mvs[i])
	}
	return movies
}

func DbUserToRest(u db.User) rest.User {
	return rest.User{
		Id:        u.ExternalId.String(),
		Username:  u.Username,
		Name:      u.Name,
		Surname:   u.Surname,
		Birhtdate: rest.DateOnly(u.Birthdate),
	}
}

func RestMovieToDb(m rest.Movie) (*db.Movie, error) {
	exId, err := uuid.Parse(m.ID)
	if err != nil {
		return nil, err
	}
	return &db.Movie{
		ExternalId:  exId,
		Name:        m.Name,
		ReleaseDate: time.Time(m.ReleaseDate),
		MinAge:      m.MinAge,
		Description: m.Description,
		Rate:        m.Rate,
		Url:         m.Url,
	}, nil
}

func RestMoviesToDb(movs []rest.Movie) ([]db.Movie, error) {
	dbMovs := make([]db.Movie, len(movs))
	for i := 0; i < len(movs); i++ {
		mov, err := RestMovieToDb(movs[i])
		if err != nil {
			return nil, err
		}
		dbMovs[i] = *mov
	}
	return dbMovs, nil
}

func RestUserToDb(u rest.User) (*db.User, error) {
	exId, err := uuid.Parse(u.Id)
	if err != nil {
		return nil, err
	}
	return &db.User{
		ExternalId: exId,
		Username:   u.Username,
		Name:       u.Name,
		Surname:    u.Surname,
		Birthdate:  time.Time(u.Birhtdate),
	}, nil
}
