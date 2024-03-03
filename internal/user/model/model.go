package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (b *Users) TableName() string {
	return "user"
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	u.Id = newUuid.String()
	return nil
}
