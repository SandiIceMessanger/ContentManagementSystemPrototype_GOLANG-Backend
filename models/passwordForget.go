package models

import (
	"github.com/jinzhu/gorm"
)

type PasswordForget struct {
	gorm.Model
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
