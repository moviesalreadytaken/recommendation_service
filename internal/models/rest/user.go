package rest

type User struct {
	Id        string   `json:"id,omitempty"`
	Username  string   `json:"username,omitempty"`
	Name      string   `json:"name,omitempty"`
	Surname   string   `json:"surname,omitempty"`
	Birhtdate DateOnly `json:"birthdate,omitempty"`
}

