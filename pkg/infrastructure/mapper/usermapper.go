package mapper

import (
	"github.com/pangud/pangud/pkg/infrastructure/mapper/po"
	"gorm.io/gorm"
)

type UserMapper struct {
	db *gorm.DB
}

func NewUserMapper(db *gorm.DB) *UserMapper {
	return &UserMapper{
		db: db,
	}
}

func (m UserMapper) ListAll([]po.UserPO, error) {

}
