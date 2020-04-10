package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ChallengeFile struct {
	Base `valid:"required"`
	Name string `json:"name" valid:"notnull"`
	URL  string `json:"name" valid:"url"`
}

func NewChallengeFile(name string, url string) (*ChallengeFile, error) {

	cfile := ChallengeFile{
		Name: name,
		URL:  url,
	}

	cfile.ID = uuid.NewV4().String()
	cfile.CreatedAt = time.Now()

	err := cfile.isValid()

	if err != nil {
		return nil, err
	}

	return &cfile, nil
}

func (cfile *ChallengeFile) isValid() (error) {

	_, err := govalidator.ValidateStruct(cfile)

	if err != nil {
		return err
	}

	return nil
}
