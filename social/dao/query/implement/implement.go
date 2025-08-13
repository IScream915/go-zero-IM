package implement

import (
	_interface "go-zero-IM/social/dao/query/interface"

	"gorm.io/gorm"
)

type DbToolHelper[T any] struct {
	DB *gorm.DB
}

func NewDbToolHelper[T any](db *gorm.DB) _interface.CommonQuery[T] {
	return &DbToolHelper[T]{DB: db}
}
