package models

import "github.com/icbd/gohighlights/utils"

func FakeUser() *User {
	email := utils.GenerateToken(10) + "@gmail.com"
	password := "12345678"
	u := User{Email: email, Password: password}
	u.CalcPasswordHash()
	DB().Create(&u)

	return &u
}
