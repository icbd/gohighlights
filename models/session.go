package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type SessionVO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"min=6"`
}

type Session struct {
	BaseModel
	Token     string    `json:"token" gorm:"uniqueindex;not null;type:varchar(255)"`
	ExpiredAt time.Time `json:"expired_at" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      *User     `json:"user,omitempty"`
}

func SessionFindByToken(token string) (session *Session, err error) {
	if token == "" {
		return nil, gorm.ErrRecordNotFound
	}

	session = &Session{}
	if err := DB().Where("token = ?", token).Preload("User").First(session).Error; err != nil {
		return nil, err
	}

	if session.ExpiredAt.Before(time.Now()) {
		return nil, fmt.Errorf("token expired")
	}

	return session, nil
}
