package models

import (
	"github.com/jinzhu/gorm"
	"github.com/go-ozzo/ozzo-validation"
	"regexp"
)

type Map struct {
	gorm.Model
	AreaName        string `json:"area_name"`
	AreaCoordinates string `json:"area_coordinates"`
	EventId         int
}

func (Map) TableName() string {
	return "rw_maps"
}

func (m Map) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.AreaName, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]+"))),
		validation.Field(&m.AreaCoordinates, validation.Required),
		validation.Field(&m.EventId,validation.Required),
	)
}
