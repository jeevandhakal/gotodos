package db

import (
	"github.com/jeevandhakal/todos/graph/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(model.User{}, model.Todo{})
	DbExceptionHandle(err)
}
