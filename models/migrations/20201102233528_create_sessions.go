package migrations

import (
	"github.com/icbd/gohighlights/models"
	"time"
)

func createSessions() error {
	type Session struct {
		models.BaseModel
		Token     string    `json:"token" gorm:"type:varchar(255);not null"`
		ExpiredAt time.Time `json:"expired_at" gorm:"not null"`
		UserID    uint      `json:"user_id" gorm:"not null"`
	}
	return mm.ChangeTable(&Session{})
}
