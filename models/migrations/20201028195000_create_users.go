package migrations

import "github.com/icbd/gohighlights/models"

func createUsers() error {
	type User struct {
		models.BaseModel
		Status       string `gorm:"type:varchar(255);index"`
		Email        string `gorm:"type:varchar(255)"`
		PasswordHash string `gorm:"type:text;comment:BCrypt"`
	}
	return mm.ChangeTable(&User{})
}
