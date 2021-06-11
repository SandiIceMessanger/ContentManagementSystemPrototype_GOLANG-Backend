package models

import (
	"github.com/jinzhu/gorm"
)

type PermissionMaster struct {
	gorm.Model
	Email  string `json:"email" form:"email"`
	Status bool   `json:"status" form:"status"`
}
