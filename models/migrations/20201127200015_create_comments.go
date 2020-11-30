package migrations

import (
	"github.com/icbd/gohighlights/models"
)

func createComments() error {
	type Comment struct {
		models.BaseModel
		UserID  uint   `gorm:"not null"`
		MarkID  uint   `gorm:"not null"`
		Content string `gorm:"type:text"`
	}
	return mm.ChangeTable(&Comment{})
}
