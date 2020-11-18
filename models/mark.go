package models

type Mark struct {
	BaseModel

	UserId    uint   `gorm:"index:idx_marks_user_id_hash_key,priority:1;not null" json:"user_id"`
	URL       string `gorm:"type:varchar(255);index" json:"url" binding:"required"`
	Tag       string `gorm:"type:varchar(255);comment:color or other tag;not null" json:"tag" binding:"required"`
	HashKey   string `gorm:"type:varchar(255);index:idx_marks_user_id_hash_key,priority:1;not null" json:"hash_key" binding:"required"`
	Selection string `gorm:"type:text;not null" json:"selection" binding:"required"`
}

type MarkCreateVO struct {
	URL       string `json:"url"`
	Tag       string `json:"tag"`
	HashKey   string `json:"hash_key"`
	Selection string `json:"selection"`
}

type MarkUpdateVO struct {
	Tag string `json:"tag"`
}
