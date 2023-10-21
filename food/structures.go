package food

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID            uuid.UUID  `json:"ID"`
	Username      string     `json:"username"`
	Password      string     `json:"-"`
	Email         string     `json:"email"`
	FavouriteFood []*Recepie `json:"foods"`
}

type Recepie struct {
	ID         uint
	Name       string
	Users      []User
	Produces   []*Produce
	Appliances []*Appliance
}

type Produce struct {
	ID       uint
	Name     string
	Recepies []*Recepie
}

type Appliance struct {
	ID       uint
	Name     string
	Recepies []*Recepie
}
