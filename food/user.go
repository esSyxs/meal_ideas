package food

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	user    map[string]*User
	userMux sync.Mutex
)

func init() {
	u1 := &User{
		ID:            uuid.NewV1(),
		Username:      "Janis",
		Password:      "",
		Email:         "janis@gmail.com",
		FavouriteFood: nil,
	}

	u2 := &User{
		ID:            uuid.NewV1(),
		Username:      "Peteris",
		Password:      "",
		Email:         "peteris@example.com",
		FavouriteFood: nil,
	}

	err := u1.HashPassword("password1")
	if err != nil {
		panic(err)
	}

	err = u2.HashPassword("password2")
	if err != nil {
		panic(err)
	}

	user = make(map[string]*User)

	user[u1.Email] = u1
	user[u2.Email] = u2
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func AddUser(u *User) error {
	userMux.Lock()
	defer userMux.Unlock()

	_, ok := user[u.Email]
	if ok {
		return errors.New("user already exists")
	}

	user[u.Email] = u

	return nil
}

func AddUserRecipe(email string, rec Recepie) error {
	userMux.Lock()
	defer userMux.Unlock()

	_, ok := user[email]
	if !ok {
		return errors.New("user not found")
	}

	user[email].FavouriteFood = append(user[email].FavouriteFood, &rec)

	return nil
}

func RemoveUserRecipe(email string, rec Recepie) error {
	userMux.Lock()
	defer userMux.Unlock()

	_, ok := user[email]
	if !ok {
		return errors.New("user not found")
	}

	for i, f := range user[email].FavouriteFood {
		if f.ID == rec.ID {
			user[email].FavouriteFood = append(user[email].FavouriteFood[:i], user[email].FavouriteFood[i+1:]...)
			return nil
		}
	}

	return nil
}

func UpdateUser(u *User, oldEmail string) error {
	userMux.Lock()
	defer userMux.Unlock()

	_, ok := user[u.Email]
	if ok && u.Email != oldEmail {
		return errors.New("user already exists")
	}

	user[u.Email] = u
	if u.Email != oldEmail {
		delete(user, oldEmail)
	}

	return nil
}

func GetUser(email string) (*User, error) {
	userMux.Lock()
	defer userMux.Unlock()

	_, ok := user[email]
	if !ok {
		return nil, errors.New("incorrect credentials")
	}

	return user[email], nil
}
