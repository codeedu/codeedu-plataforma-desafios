package domain

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

// Base contains common columns for all tables.
type Base struct {
	ID        string `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Timer
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {

	scope.SetColumn("CreatedAt", time.Now())
	return scope.SetColumn("ID", uuid.NewV4().String())
}
