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
	Email        string `json:"email" binding:"required" gorm:"uniqueIndex"`
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

func (u *User) Validate() error {
	return _validator.Struct(u)
}

func FindUserByEmail(email string) (u *User, err error) {
	u = &User{}
	err = DB().Where("email = ?", email).First(u).Error
	return
}

func (u *User) Create() error {
	if err := u.CalcPasswordHash(); err != nil {
		return err
	}

	return DB().Create(u).Error
}

func (u *User) GenerateSession() (*Session, error) {
	s := Session{
		Token:     utils.GenerateToken(16),
		ExpiredAt: time.Now().AddDate(0, 1, 0),
		UserID:    u.ID}
	if err := DB().Create(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}
