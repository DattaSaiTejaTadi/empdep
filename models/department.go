package models

type Department struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Floor int    `json:"floor,omitempty"`
}
