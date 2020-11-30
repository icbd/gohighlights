package models

type Comment struct {
	BaseModel
	UserID  uint   `gorm:"not null" json:"-" binding:"required"`
	User    *User  `json:"user,omitempty"`
	MarkID  uint   `gorm:"index;not null" json:"mark_id" binding:"required"`
	Content string `gorm:"type:text" json:"content" binding:"required"`
}

type CommentVO struct {
	Content string `json:"content" binding:"required"`
}

func CommentFind(commentID uint) (comment *Comment, err error) {
	comment = &Comment{}
	err = DB().First(&comment, commentID).Error
	return comment, err
}

func CommentFindByMarkID(userID uint, markID uint) (comment *Comment, err error) {
	comment = &Comment{}
	err = DB().Where("user_id = ? AND mark_id = ?", userID, markID).First(&comment).Error
	return comment, err
}

func CommentCreate(userID uint, markID uint, content string) (comment *Comment, err error) {
	comment = &Comment{
		UserID:  userID,
		MarkID:  markID,
		Content: content,
	}
	err = DB().Create(comment).Error
	return comment, err
}

func CommentUpdate(userID uint, markID uint, content string) (comment *Comment, err error) {
	comment, err = CommentFindByMarkID(userID, markID)
	if err != nil {
		return nil, err
	}
	err = DB().Model(&comment).Update("content", content).Error
	return comment, err
}

func CommentDestroy(userID uint, markID uint) (comment *Comment, err error) {
	comment, err = CommentFindByMarkID(userID, markID)
	if err != nil {
		return nil, err
	}
	err = DB().Delete(comment).Error
	return comment, err
}
