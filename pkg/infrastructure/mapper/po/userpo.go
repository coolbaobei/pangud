package po

import (
	"gorm.io/gorm"
)

type UserPO struct {
	gorm.Model
	Name string
}
