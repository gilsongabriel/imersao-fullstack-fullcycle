package models

import (
	validator "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Base     `valid:"required"`
	Name     string     `json:"name" valid:"notnull"`
	Email     string     `json:"email" valid:"email"`
}

func (user *User) isValid() error {
	_, err := validator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(name string, email string) (*User, error) {
	user := User{
		Email: email,
		Name: name,
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}