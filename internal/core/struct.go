package core

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Struct struct {
	ID        uint `gorm:"primary_key"`
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *Struct) BeforeCreate(tx *gorm.DB) (err error) {
	switch {
	case u.Message == "":
		return errors.New("message cannot be empty")
	default:
		return nil
	}
}
