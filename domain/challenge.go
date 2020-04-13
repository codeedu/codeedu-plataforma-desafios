package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Challenge struct {
	Base           `valid:"required"`
	Title          string          `json:"title" valid:"notnull" gorm:"type:varchar(255)"`
	Slug           string          `json:"slug" valid:"notnull" gorm:"type:varchar(255)"`
	Description    string          `json:"description" valid:"notnull" gorm:"type:text"`
	Tags           string          `json:"tags" valid:"notnull" gorm:"type:varchar(255)"`
	Requirements   string          `json:"requirements" valid:"notnull" gorm:"type:varchar(255)"`
	Level          string          `json:"level" valid:"notnull" gorm:"type:varchar(255)"`
	ChallengeFiles []*ChallengeFile `gorm:"ForeignKey:ChallengeID" valid:"-"`
}

func NewChallenge() *Challenge {
	return &Challenge{}
}

func (challenge *Challenge) Valid() error {

	challenge.ID = uuid.NewV4().String()
	challenge.CreatedAt = time.Now()

	_, err := govalidator.ValidateStruct(challenge)

	if err != nil {
		return err
	}

	return nil

}