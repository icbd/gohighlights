package migrations

import (
	"gorm.io/gorm"
)

func createUsers() error {
	type User struct {
		gorm.Model
		Status       uint8  `gorm:"type:tinyint;comment:enum"`
		Email        string `gorm:"type:varchar(255);uniqueIndex"`
		PasswordHash string `gorm:"type:text;comment:BCrypt"`
	}
	return mm.ChangeTable(&User{})
}
