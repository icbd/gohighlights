package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt DeletedAt `gorm:"default:0;index" json:"-"`
}

func DB() *gorm.DB {
	return db
}

func Ping() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Ping(); err != nil {
		return err
	}

	return nil
}

type PaginationVO struct {
	Page int `json:"page" form:"page" binding:"min=1"`
	Size int `json:"size" form:"size" binding:"max=100"`
}

var Pagination = PaginationVO{Page: 1, Size: 10}

func PaginationScope(vo PaginationVO) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if err := _validator.Struct(&vo); err != nil {
			return db.Limit(10).Offset(0)
		}
		return db.Limit(vo.Size).Offset(vo.Size * (vo.Page - 1))
	}
}
