package models

import (
	"github.com/jinzhu/gorm"
)

type PasswordForget struct {
	gorm.Model
	Token  string `json:"token" form:"token"`
	IdUser int    `json:"id_user" form:"id_user"`
	// User   User   `gorm:"foreignKey:IdUser" json:"user" form:"user"`
}
