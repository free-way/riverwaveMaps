package models

import (
	"github.com/jinzhu/gorm"
)
type Map struct {
	gorm.Model
	AreaName string
	EventId int
}

func (Map) TableName() string  {
	return "rw_maps"
}