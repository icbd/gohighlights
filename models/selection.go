package models

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/spf13/cast"
)

type Selection struct {
	Texts       []string `json:"texts" binding:"required"`
	StartOffset int      `json:"startOffset" binding:"required"`
	EndOffset   int      `json:"endOffset" binding:"required"`
}

// Scan from DB
func (s *Selection) Scan(value interface{}) error {
	v := cast.ToString(value)
	return json.Unmarshal([]byte(v), s)
}

// Save to DB
func (s Selection) Value() (driver.Value, error) {
	if j, err := json.Marshal(s); err != nil {
		return nil, err
	} else {
		return string(j), nil
	}
}
