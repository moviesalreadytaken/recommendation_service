package serivces

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/moviesalreadytaken/recommendation_service/internal/models/rest"
	"github.com/moviesalreadytaken/recommendation_service/internal/utils"
)

type MoviesServiceClient struct {
	serviceUrl string
	resty      *resty.Client
}

func NewMovieServiceClient(cnf *utils.AppConfig) (*MoviesServiceClient, error) {
	_, err := url.Parse(cnf.MoviesServiceUrl)
	if err != nil {
		return nil, err
	}
	return &MoviesServiceClient{
		serviceUrl: cnf.MoviesServiceUrl,
		resty:      resty.New(),
	}, nil
}

func (c *MoviesServiceClient) GetAllMovies() ([]rest.Movie, error) {
	resp, err := c.resty.R().Get(fmt.Sprintf("%s/movies", c.serviceUrl))
	if err != nil {
		return nil, err
	}
	var movies []rest.Movie
	err = json.Unmarshal(resp.Body(), &movies)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
