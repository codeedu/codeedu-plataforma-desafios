package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Author struct {
	Base    `valid:"required"`
	Name    string `json:"name" valid:"notnull" gorm:"type:varchar(255)"`
	Email   string `json:"email" valid:"email" gorm:"type:varchar(255)"`
	Picture string `json:"picture" valid:"-" gorm:"type:varchar(255)"`
	Github  string `json:"picture" valid:"-" gorm:"type:varchar(255)"`
}

func NewAuthor(name string, email string) (*Author, error) {

	author := Author{
		Name:  name,
		Email: email,
	}

	author.ID = uuid.NewV4().String()
	author.CreatedAt = time.Now()

	err := author.isValid()

	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (author *Author) isValid() (error) {

	_, err := govalidator.ValidateStruct(author)

	if err != nil {
		return err
	}

	return nil
}
