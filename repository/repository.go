package repository

import (
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) template.Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (r repository) TemplateGet() ([]template.Template, error) {
	panic("implement me")
}
