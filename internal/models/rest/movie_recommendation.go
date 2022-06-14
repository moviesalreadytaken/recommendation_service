package rest

import "time"

type MovieRecommendation struct {
	ID                 string    `json:"id,omitempty"`
	Movies             []Movie   `json:"movies,omitempty"`
	User               User      `json:"user,omitempty"`
	RecommendationDate time.Time `json:"recommendationDate,omitempty"`
	Description        string    `json:"description,omitempty"`
}
