package models

import "time"

type Person struct {
	ID        string     `json:"id"`
	Firstname string     `json:"firstname"`
	Lastname  string     `json:"lastname"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type PersonCreateModel struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type PersonUpdateModel struct {
	ID        string `json:"-"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
