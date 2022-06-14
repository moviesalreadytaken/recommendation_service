package rest

type Movie struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	ReleaseDate DateOnly `json:"releaseDate,omitempty"`
	MinAge      uint32   `json:"minAge,omitempty"`
	Description string   `json:"description,omitempty"`
	Rate        float32  `json:"rate,omitempty"`
	Url         string   `json:"url,omitempty"`
}
