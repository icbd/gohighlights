package models

import (
	"fmt"
	"time"
)

type SessionVO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"min=6"`
}

func (s *SessionVO) Validate() error {
	return _validator.Struct(s)
}

func (s *SessionVO) CurrentUser() (u *User, ok bool) {
	u = &User{Email: s.Email, Password: s.Password}
	DB().Where("email = ?", u.Email).First(u)
	if DB().Error != nil {
		return nil, false
	}

	return u, u.ValidPassword()
}

type Session struct {
	BaseModel
	Token     string    `json:"token" gorm:"uniqueindex;not null;type:varchar(255)"`
	ExpiredAt time.Time `json:"expired_at" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user"`
}

func (s *Session) FindByToken() error {
	if err := DB().Where("token = ?", s.Token).Preload("User").First(s).Error; err != nil {
		return err
	}

	if s.ExpiredAt.Before(time.Now()) {
		return fmt.Errorf("token expired")
	}

	return nil
}
