package models

import (
	"encoding/base64"
	"net/url"
	"strings"
)

type Mark struct {
	BaseModel

	UserId    uint      `gorm:"index:idx_marks_user_id_hash_key,priority:1;not null" json:"user_id"`
	URL       string    `gorm:"type:varchar(255);index" json:"url" binding:"required"`
	Tag       string    `gorm:"type:varchar(255);comment:color or other tag;not null" json:"tag" binding:"required"`
	HashKey   string    `gorm:"type:varchar(255);index:idx_marks_user_id_hash_key,priority:1;not null" json:"hash_key" binding:"required"`
	Selection Selection `gorm:"type:text;not null" json:"selection" binding:"required"`
}

type MarkCreateVO struct {
	URL       string    `json:"url"`
	Tag       string    `json:"tag"`
	HashKey   string    `json:"hash_key"`
	Selection Selection `json:"selection"`
}

type MarkUpdateVO struct {
	HashKey   string    `json:"-"`
	Tag       string    `json:"tag"`
	Selection Selection `json:"selection"`
}

func MarkFetchBy(u *User, hashKey string) (mark *Mark, err error) {
	mark = &Mark{}
	err = DB().Where("user_id = ? AND hash_key = ?", u.ID, hashKey).First(mark).Error
	return
}

func MarkQuery(u *User, safeBase64URL string) []Mark {
	marks := make([]Mark, 0)
	var err error
	if safeBase64URL, err = url.QueryUnescape(safeBase64URL); err == nil {
		if urlBytes, err := base64.StdEncoding.DecodeString(safeBase64URL); err == nil {
			DB().Where("user_id = ? AND url = ?", u.ID, string(urlBytes)).Find(&marks)
		}
	}
	return marks
}

func MarkCreate(u *User, vo MarkCreateVO) (mark *Mark, err error) {
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

func MarkUpdate(u *User, vo *MarkUpdateVO) (mark *Mark, err error) {
	mark, err = MarkFetchBy(u, vo.HashKey)
	if err != nil {
		return nil, err
	}
	mark.Tag = vo.Tag
	mark.Selection = vo.Selection
	err = DB().Updates(mark).Error
	return mark, err
}

func MarkDestroy(u *User, hashKey string) (mark *Mark, err error) {
	mark = &Mark{}
	if err := DB().Where("user_id = ? AND hash_key = ?", u.ID, hashKey).First(&mark).Error; err != nil {
		return nil, err
	}
	err = DB().Delete(&mark).Error
	return mark, err
}

func MarkList(u *User, vo PaginationVO) []Mark {
	var marks []Mark
	DB().Scopes(PaginationScope(vo)).Where("user_id = ?", u.ID).Find(&marks)
	return marks
}

func MarkTotal(u *User) (total int64) {
	DB().Model(&Mark{}).Where("user_id = ?", u.ID).Count(&total)
	return total
}

func (m *Mark) SelectionText() string {
	texts := m.Selection.Texts
	if len(texts) == 1 {
		return texts[0][m.Selection.StartOffset:m.Selection.EndOffset]
	}

	texts[0] = texts[0][m.Selection.StartOffset:len(texts[0])]
	lastIndex := len(texts) - 1
	texts[lastIndex] = texts[lastIndex][0:m.Selection.EndOffset]
	return strings.Join(texts, "\r\n")
}
