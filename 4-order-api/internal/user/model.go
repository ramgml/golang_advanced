package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone  string        `gorm:"uniqueIndex"`
}
