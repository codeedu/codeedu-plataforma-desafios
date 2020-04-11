package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ChallengeFile struct {
	Base        `valid:"required"`
	Name        string    `json:"name" valid:"notnull" gorm:"type:varchar(255)"`
	URL         string    `json:"name" valid:"url" gorm:"type:varchar(255)"`
	ChallengeID string    `gorm:"column:challenge_id;type:uuid;not null" valid:"-"`
	Challenge   Challenge `valid:"-"`
}

func NewChallengeFile(name string, url string) (*ChallengeFile, error) {

	cfile := ChallengeFile{
		Name: name,
		URL:  url,
	}

	cfile.ID = uuid.NewV4().String()
	cfile.CreatedAt = time.Now()

	err := cfile.Valid()

	if err != nil {
		return nil, err
	}

	return &cfile, nil
}

func (cfile *ChallengeFile) Valid() (error) {

	_, err := govalidator.ValidateStruct(cfile)

	if err != nil {
		return err
	}

	return nil
}
