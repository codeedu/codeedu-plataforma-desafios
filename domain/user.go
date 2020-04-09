package domain

import (
	"errors"
	"github.com/badoux/checkmail"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Base
	Name     string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `json:"-"`
	Token    string `json:"token" gorm:"unique_index"`
}

func NewUser() *User {
	return &User{}
}

func (user *User) Prepare() error {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Error during the password generation: %v", err)
	}

	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	err = user.validate()

	if err != nil {
		log.Fatalf("Error during the user validation: %v", err)
	}

	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("name cannot be empty")
	}
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid email")
	}
	if user.Password == "" {
		return errors.New("pasword cannot be empty")
	}

	return nil
}
