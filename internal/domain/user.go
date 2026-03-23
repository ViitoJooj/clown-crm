package domain

import (
	"errors"
	"regexp"
	"time"

	"github.com/ViitoJooj/clown-crm/pkg/cryptography"
	"github.com/google/uuid"
)

type User struct {
	UUID       string
	First_Name string
	Last_Name  string
	Email      string
	Password   string
	Updated_at time.Time
	Created_at time.Time
}

func NewUser(first_name string, last_name string, email string, password string) (*User, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, errors.New("error on gen uuid: " + uuid.String())
	}

	if !nameIsValid(first_name) {
		return nil, errors.New("name is not valid")
	}

	if !nameIsValid(last_name) {
		return nil, errors.New("name is not valid")
	}

	if !emailIsValid(email) {
		return nil, errors.New("email is not valid")
	}

	if !passwordIsValid(password) {
		return nil, errors.New("password is not valid")
	}

	hashPassword, err := cryptography.HashPassword(password)
	if err != nil {
		return nil, errors.New("error on hash password")
	}

	user := User{
		UUID:       uuid.String(),
		First_Name: first_name,
		Last_Name:  last_name,
		Email:      email,
		Password:   hashPassword,
		Updated_at: time.Now(),
		Created_at: time.Now(),
	}

	return &user, nil
}

func nameIsValid(name string) bool {
	hasInvalid := regexp.MustCompile(`[^a-zA-ZÀ-ÿ0-9\s-]+`)

	if hasInvalid.MatchString(name) {
		return false
	}

	if len(name) > 100 || len(name) < 3 {
		return false
	}

	return true
}

func emailIsValid(email string) bool {
	hasInvalid := regexp.MustCompile(`^[a-zA-Z0-9._%+\\-]+@[a-zA-Z0-9.\\-]+\\.[a-zA-Z]{2,4}$`)

	if hasInvalid.MatchString(email) {
		return false
	}

	if len(email) > 100 || len(email) < 3 {
		return false
	}

	return true
}

func passwordIsValid(password string) bool {
	if len(password) > 100 || len(password) < 6 {
		return false
	}

	return true
}
