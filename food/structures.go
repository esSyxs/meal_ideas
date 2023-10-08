package food

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID            uuid.UUID
	Username      string
	Password      string
	Email         string
	FavouriteFood []*Recepies
}

type Recepies struct {
	ID         uint
	Name       string
	Users      []User
	Produces   []*Produce
	Appliances []*Appliances
}

type Produce struct {
	ID       uint
	Name     string
	Recepies []*Recepies
}

type Appliances struct {
	ID       uint
	Name     string
	Recepies []*Recepies
}
