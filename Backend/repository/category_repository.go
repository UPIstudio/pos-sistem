package repository

import (
	"github.com/luthfi/pos/config"
	"github.com/luthfi/pos/models"
)

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	Create(c *models.Category) error
	Update(id uint, c *models.Category) error
	Delete(id uint) error
}

type categoryRepo struct{}

func NewCategoryRepository() CategoryRepository { return &categoryRepo{} }

func (r *categoryRepo) GetAll() (cat []models.Category, err error) {
	err = config.DB.Find(&cat).Error
	return
}

func (r *categoryRepo) Create(c *models.Category) error {
	return config.DB.Create(c).Error
}

func (r *categoryRepo) Update(id uint, c *models.Category) error {
	return config.DB.Model(&models.Category{}).Where("id = ?", id).Updates(c).Error
}

func (r *categoryRepo) Delete(id uint) error {
	return config.DB.Delete(&models.Category{}, id).Error
}
