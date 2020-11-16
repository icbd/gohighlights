package models

import (
	"fmt"
	"gorm.io/gorm"
)

func RegisterGormCallBack() (err error) {
	if err = db.Callback().Create().Before("gorm:begin_transaction").Register("uniqCheck", uniqCheck); err != nil {
		return err
	}
	if err = db.Callback().Update().Before("gorm:begin_transaction").Register("uniqCheck", uniqCheck); err != nil {
		return err
	}

	return nil
}

// uniqCheck is a callback function for Gorm.
// Run query in a separate transaction before create/update transaction.
// Final uniqueness is limited by the unique index of the database.
func uniqCheck(db *gorm.DB) {
	schema := db.Statement.Schema
	for _, field := range schema.Fields {
		if field.TagSettings["UNIQUEINDEX"] == "UNIQUEINDEX" {
			var total int64
			fieldValue, _ := field.ValueOf(db.Statement.ReflectValue)
			err := db.Transaction(func(tx *gorm.DB) error {
				tx.Table(schema.Table).Where(field.DBName+" = ?", fieldValue).Count(&total)
				return tx.Error
			})
			if err != nil {
				_ = db.AddError(err)
				return
			}
			if total != 0 {
				_ = db.AddError(fmt.Errorf(field.Name + " should be uniqueness"))
				return
			}
		}
	}
}
