package models

import (
	"encoding/base64"
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

func (u *User) MarkQuery(url string) []Mark {
	var marks []Mark
	var urlStr string
	if urlBytes, err := base64.StdEncoding.DecodeString(url); err != nil {
		urlStr = url
	} else {
		urlStr = string(urlBytes)
	}
	DB().Where("user_id = ? AND url = ?", u.ID, urlStr).Find(&marks)
	return marks
}

func (u *User) MarksAll(vo PaginationVO) []Mark {
	var marks []Mark
	DB().Scopes(PaginationScope(vo)).Where("user_id = ?", u.ID).Find(&marks)
	return marks
}

func (u *User) MarksTotal() (total int64) {
	DB().Model(&Mark{}).Where("user_id = ?", u.ID).Count(&total)
	return total
}

func (u *User) CreateMark(vo MarkCreateVO) (mark *Mark, err error) {
	mark = &Mark{
		UserId:    u.ID,
		URL:       vo.URL,
		Tag:       vo.Tag,
		HashKey:   vo.HashKey,
		Selection: vo.Selection,
	}
	if err := DB().Create(mark).Error; err != nil {
		return nil, err
	} else {
		return mark, nil
	}
}
