package serivces

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/moviesalreadytaken/recommendation_service/internal/models/rest"
	"github.com/moviesalreadytaken/recommendation_service/internal/utils"
)

type UsersServiceClient struct {
	serviceUrl string
	resty      *resty.Client
}

var ErrUserNotFound = errors.New("user not found")

func NewUsersServiceClient(cnf *utils.AppConfig) (*UsersServiceClient, error) {
	_, err := url.Parse(cnf.UsersServiceUrl)
	if err != nil {
		return nil, err
	}
	return &UsersServiceClient{
		resty:      resty.New(),
		serviceUrl: cnf.UsersServiceUrl,
	}, nil
}

func (c *UsersServiceClient) GetUserById(userId uuid.UUID) (*rest.User, error) {
	resp, err := c.resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(rest.UserByIdRequest{
			Id: userId.String(),
		}).
		Post(fmt.Sprintf("%s/users/find", c.serviceUrl))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, ErrUserNotFound
	}
	var user rest.User
	err = json.Unmarshal(resp.Body(), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
