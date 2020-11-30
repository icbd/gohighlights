package models

import (
	"bytes"
	"encoding/base64"
	"gorm.io/gorm"
	"net/url"
	"strings"
)

type Mark struct {
	BaseModel
	UserID    uint      `gorm:"not null" json:"user_id" binding:"required"`
	URL       string    `gorm:"type:varchar(255);index" json:"url" binding:"required"`
	Tag       string    `gorm:"type:varchar(255);comment:color or other tag;not null" json:"tag" binding:"required"`
	HashKey   string    `gorm:"type:varchar(255);not null" json:"hash_key" binding:"required"`
	Selection Selection `gorm:"type:text;not null" json:"selection" binding:"required"`

	Comment *Comment `json:"comment,omitempty"` // has one
}

type MarkCreateVO struct {
	URL       string    `json:"url" binding:"required"`
	Tag       string    `json:"tag" binding:"required"`
	HashKey   string    `json:"hash_key" binding:"required"`
	Selection Selection `json:"selection" binding:"required"`
}

type MarkUpdateVO struct {
	Tag       string    `json:"tag"`
	//Selection Selection `json:"selection"`
}

func MarkFindByHashKey(userID uint, hashKey string) (mark *Mark, err error) {
	if hashKey == "" {
		return nil, gorm.ErrRecordNotFound
	}
	mark = &Mark{}
	err = DB().Preload("Comment").Where("user_id = ? AND hash_key = ?", userID, hashKey).First(mark).Error
	return
}

func MarkFind(markID uint) (mark *Mark, err error) {
	mark = &Mark{}
	err = DB().Preload("Comment").First(mark, markID).Error
	return mark, err
}

func MarkQuery(userID uint, safeBase64URL string) []*Mark {
	marks := make([]*Mark, 0)
	var err error
	if safeBase64URL, err = url.QueryUnescape(safeBase64URL); err == nil {
		if urlBytes, err := base64.StdEncoding.DecodeString(safeBase64URL); err == nil {
			DB().Preload("Comment").Where("user_id = ? AND url = ?", userID, string(urlBytes)).Find(&marks)
		}
	}
	return marks
}

func MarkCreate(userID uint, vo *MarkCreateVO) (mark *Mark, err error) {
	mark = &Mark{
		UserID:    userID,
		URL:       vo.URL,
		Tag:       vo.Tag,
		HashKey:   vo.HashKey,
		Selection: vo.Selection,
	}
	err = DB().Create(mark).Error
	return mark, err
}

func MarkUpdate(userID uint, hashKey string, vo *MarkUpdateVO) (mark *Mark, err error) {
	mark, err = MarkFindByHashKey(userID, hashKey)
	if err != nil {
		return nil, err
	}
	mark.Tag = vo.Tag
	//mark.Selection = vo.Selection
	err = DB().Updates(mark).Error
	return mark, err
}

func MarkDestroy(userID uint, hashKey string) (mark *Mark, err error) {
	mark, err = MarkFindByHashKey(userID, hashKey)
	if err != nil {
		return nil, err
	}
	err = DB().Select("Comment").Delete(&mark).Error
	return mark, err
}

func MarkList(userID uint, vo PaginationVO) []Mark {
	var marks []Mark
	DB().Scopes(PaginationScope(vo)).Preload("Comment").Where("user_id = ?", userID).Find(&marks)
	return marks
}

func MarkTotal(userID uint) (total int64) {
	DB().Model(&Mark{}).Where("user_id = ?", userID).Count(&total)
	return total
}

const HTMLSplitTag = "\t"

func (m *Mark) SelectionText() string {
	texts := m.Selection.Texts
	if len(texts) == 1 {
		runes := []rune(texts[0])
		subRunes := string(runes[m.Selection.StartOffset:m.Selection.EndOffset])
		return strings.TrimSpace(subRunes)
	}

	firstUnicodeStr := []rune(texts[0])
	texts[0] = string(firstUnicodeStr[m.Selection.StartOffset:len(firstUnicodeStr)])

	lastIndex := len(texts) - 1
	lastUnicodeStr := []rune(texts[lastIndex])
	texts[lastIndex] = string(lastUnicodeStr[0:m.Selection.EndOffset])

	var buffer bytes.Buffer
	for i, s := range texts {
		if i > 0 {
			buffer.WriteString(HTMLSplitTag)
		}
		buffer.WriteString(strings.TrimSpace(s))
	}
	return buffer.String()
}
