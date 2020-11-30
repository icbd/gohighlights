package migrations

import (
	"github.com/icbd/gohighlights/models"
)

func createMarks() error {
	type Mark struct {
		models.BaseModel
		UserID    uint   `gorm:"not null"`
		URL       string `gorm:"type:varchar(255);index;not null"`
		Tag       string `gorm:"type:varchar(255);comment:color or other tag;not null"`
		HashKey   string `gorm:"type:varchar(255);not null"`
		Selection string `gorm:"type:text;not null"`
	}
	return mm.ChangeTable(&Mark{})
}
