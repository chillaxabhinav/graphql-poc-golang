package myModels

import (
	"graphql-poc/graph/model"
)

type User struct {
	Firstname string         `json:"firstname"`
	Lastname  string         `json:"lastname"`
	Address   *model.Address `json:"address"`
}
