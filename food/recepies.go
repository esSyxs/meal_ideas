package food

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	recID       autoIncrement
	recepies    map[uint]*Recepie
	recepiesMux sync.Mutex
	aplID       autoIncrement
	appliances  map[string]*Appliance
	produceID   autoIncrement
	produce     map[string]*Produce
)

func init() {
	recepies = make(map[uint]*Recepie)
	appliances = make(map[string]*Appliance)
	produce = make(map[string]*Produce)

	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("Produce %d", i)
		produce[name] = &Produce{
			ID:       uint(produceID.id()),
			Name:     name,
			Recepies: nil,
		}
	}

	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("Appliance %d", i)
		appliances[name] = &Appliance{
			ID:       uint(aplID.id()),
			Name:     name,
			Recepies: nil,
		}
	}

	for i := 1; i <= 10; i++ {
		id := uint(recID.id())
		recepies[id] = &Recepie{
			ID:         id,
			Name:       fmt.Sprintf("Food %d", i),
			Desciption: fmt.Sprintf("Super tasty %d", i),
			Users:      nil,
			Produces:   nil,
			Appliances: nil,
		}

		// very bad practice
		app1 := randInt(1, 3)
		<-time.After(time.Millisecond)
		app2 := randInt(4, 7)
		<-time.After(time.Millisecond)
		app3 := randInt(8, 10)

		for i := 1; i <= 10; i++ {
			if i == app1 || i == app2 || i == app3 {
				recepies[id].Appliances = append(recepies[id].Appliances, appliances[fmt.Sprintf("Appliance %d", i)])
			}
		}

		// very bad practice
		<-time.After(time.Millisecond)
		app4 := randInt(1, 5)
		<-time.After(time.Millisecond)
		app5 := randInt(6, 10)

		for i := 1; i <= 10; i++ {
			if i == app4 || i == app5 {
				recepies[id].Produces = append(recepies[id].Produces, produce[fmt.Sprintf("Produce %d", i)])
			}
		}
	}
}

func GetRecepies() map[uint]*Recepie {
	return recepies
}

func GetRecepie(id uint) (*Recepie, error) {
	recepiesMux.Lock()
	defer recepiesMux.Unlock()

	_, ok := recepies[id]
	if !ok {
		return nil, errors.New("incorrect recepie name")
	}

	return recepies[id], nil
}

func randInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}
