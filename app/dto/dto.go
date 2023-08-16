package dto

import "gorm.io/gorm"

type Dto struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dto {
	return &Dto{engine: engine}
}
