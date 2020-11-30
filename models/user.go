package models

import (
	"fmt"
	"github.com/icbd/gohighlights/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	BaseModel
	Status       string `json:"status"`
	Email        string `json:"email" binding:"required"`
	Avatar       string `json:"avatar"`
	PasswordHash string `json:"-"`
	Password     string `json:"-" gorm:"-"`
}

type UserStatus string

func (s UserStatus) String() string {
	return string(s)
}

const (
	ActivatedStatus UserStatus = "ActivatedStatus"
	InactiveStatus  UserStatus = "InactiveStatus"
)

var UserStatuses = map[string]UserStatus{
	"ActivatedStatus": ActivatedStatus,
	"InactiveStatus":  InactiveStatus,
}

func (u *User) CalcPasswordHash() error {
	if len(u.Password) < 6 {
		return fmt.Errorf("password too short")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	u.PasswordHash = string(hash)
	return nil
}

func (u *User) ValidPassword() bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password)) == nil
}

func UserFindByEmail(email string) (u *User, err error) {
	u = &User{}
	err = DB().Where("email = ?", email).First(u).Error
	return
}

func UserCreate(email string, password string) (u *User, err error) {
	u = &User{Email: email, Password: password}
	if err := u.CalcPasswordHash(); err != nil {
		return nil, err
	}
	err = DB().Create(u).Error
	return u, err
}

func (u *User) GenerateSession() (s *Session, err error) {
	s = &Session{
		Token:     utils.GenerateToken(16),
		ExpiredAt: time.Now().AddDate(0, 1, 0),
		UserID:    u.ID}
	err = DB().Create(&s).Error
	s.User = u
	return s, nil
}
