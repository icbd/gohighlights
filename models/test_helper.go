package models

import (
	"github.com/icbd/gohighlights/utils"
	"math/rand"
	"time"
)

func FakeUser() *User {
	email := utils.GenerateToken(10) + "@gmail.com"
	password := "12345678"
	u := User{Email: email, Password: password}
	u.CalcPasswordHash()
	DB().Create(&u)

	return &u
}

func FakeSelection() *Selection {
	rand.Seed(time.Now().Unix())
	texts := make([]string, rand.Intn(3))
	for i, _ := range texts {
		texts[i] = utils.GenerateToken(5 + rand.Intn(10))
	}

	startOffset := 0 + rand.Intn(5)
	endOffset := 5 + rand.Intn(5)
	if len(texts) > 1 {
		startOffset = rand.Intn(len(texts[0]))
		endOffset = rand.Intn(len(texts[len(texts)-1]))
	}

	return &Selection{
		Texts:       texts,
		StartOffset: startOffset,
		EndOffset:   endOffset,
	}
}

func FakeMark(userID uint) *Mark {
	mark := &Mark{
		UserID:    userID,
		URL:       "http://localhost/",
		Tag:       "blue",
		HashKey:   utils.GenerateToken(16),
		Selection: *FakeSelection(),
	}
	DB().Create(mark)

	return mark
}

func FakeComment(userID uint, markID uint) *Comment {
	comment := &Comment{
		UserID:  userID,
		MarkID:  markID,
		Content: utils.GenerateToken(10),
	}
	DB().Create(comment)
	return comment
}
