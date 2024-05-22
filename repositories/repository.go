package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}
type LinkRepository interface {
	GetLink(ShortURL string) (models.Link, error)
	GetLongURLCheck(ShortURL string) (models.Link, error)
	CreateLink(link models.Link) (models.Link, error)
}

func RepositoryLink(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetLink(ShortURL string) (models.Link, error) {
	var link models.Link
	err := r.db.First(&link, "short_url=?", ShortURL).Error

	return link, err
}

func (r *repository) GetLongURLCheck(ShortURL string) (models.Link, error) {
	var link models.Link
	err := r.db.First(&link, "long_url=?", ShortURL).Error

	return link, err
}

func (r *repository) CreateLink(link models.Link) (models.Link, error) {
	err := r.db.Create(&link).Error

	return link, err
}
