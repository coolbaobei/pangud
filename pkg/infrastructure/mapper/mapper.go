package mapper

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	return nil
}

var MapperSet = wire.NewSet(DB, NewUserMapper)
