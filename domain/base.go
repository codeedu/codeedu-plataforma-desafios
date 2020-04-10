package domain

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

type Base struct {
	ID        string     `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error  {

	err := scope.SetColumn("CreatedAt", time.Now())

	if err != nil {
		log.Fatalf("Error during obj creating: %v", err)
	}

	err = scope.SetColumn("ID", uuid.NewV4().String())

	if err != nil {
		log.Fatalf("Error during obj creating: %v", err)
	}

	return nil
}