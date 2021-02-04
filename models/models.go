package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	UUID     uuid.UUID `json:"_id" gorm:"primaryKey;autoIncrement:false;unique"`
	CreateAt string    `json:"create_at"`
	UpdateAt string    `json:"update_at"`
}

func GenerateISOString() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.UUID = uuid.New()

	t := GenerateISOString()

	base.CreateAt, base.UpdateAt = t, t

	return nil
}

func (base *Base) AfterUpdate(tx *gorm.DB) error {
	base.UpdateAt = GenerateISOString()
	return nil
}
