package domain

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func init() {

	govalidator.CustomTypeTagMap.Set("notempty", func(i interface{}, o interface{}) bool {
		return utf8.RuneCountInString(strings.TrimSpace(fmt.Sprintf("%v", i))) > 0
	})

	govalidator.SetFieldsRequiredByDefault(true)

}

type User struct {
	Base     `valid:"required"`
	Name     string `json:"name" gorm:"type:varchar(255)" valid:"notnull,notempty"`
	Email    string `json:"email" gorm:"type:varchar(255);unique_index" valid:"notnull,email"`
	Password string `json:"-" gorm:"type:varchar(255)" valid:"notnull"`
	Token    string `json:"token" gorm:"type:varchar(255);unique_index" valid:"notnull,uuid"`
}

func isAStrongPassword(password string) bool {

	re := []*regexp.Regexp{
		regexp.MustCompile(`[a-zA-Z]+`),
		regexp.MustCompile(`[0-9]+`),
		regexp.MustCompile(`[\_\@\.\#\$\&\+\-]+`),
	}

	for _, r := range re {
		rawData := []byte(password)
		indexes := r.FindIndex(rawData)
		if len(indexes) < 2 {
			return false
		}
	}
	return true
}

func NewUser(name string, email string, password string) (*User, error) {


	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := user.Prepare()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) Prepare() error {

	if !isAStrongPassword(user.Password) {
		return fmt.Errorf("weak password")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	err = user.validate()

	if err != nil {
		return err
	}

	return nil

}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) validate() error {

	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}

	return nil
}
