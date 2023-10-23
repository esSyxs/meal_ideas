package food

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID            uuid.UUID  `json:"ID"`
	Username      string     `json:"username" binding:"required"`
	Password      string     `json:"password" binding:"required"`
	Email         string     `json:"email" binding:"required"`
	FavouriteFood []*Recepie `json:"foods"`
}

type Recepie struct {
	ID         uint
	Name       string
	Desciption string
	Users      []User
	Produces   []*Produce
	Appliances []*Appliance
}

type Produce struct {
	ID       uint
	Name     string
	Recepies []*Recepie
}

type autoIncrement struct {
	sync.Mutex
	i int
}

type Appliance struct {
	ID   uint
	Name string
	// Not sure this filed is actually needed!
	Recepies []*Recepie
}

func (a *autoIncrement) id() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.i + 1
	a.i++

	return
}
