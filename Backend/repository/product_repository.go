package repository

import (
	"github.com/luthfi/pos/config"
	"github.com/luthfi/pos/models"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	Create(p *models.Product) error
	Update(id int, p *models.Product) error
	Delete(id int) error
	Search(query string) ([]models.Product, error)
}

type productRepo struct{}

func NewProductRepository() ProductRepository {
	return &productRepo{}
}

func (r *productRepo) GetAll() (p []models.Product, err error) {
	err = config.DB.Find(&p).Error
	return
}

func (r *productRepo) Search(query string) (p []models.Product, err error) {
	err = config.DB.Where("name ILIKE ? OR category ILIKE ?", "%"+query+"%", "%"+query+"%").Find(&p).Error
	return
}

func (r *productRepo) Create(p *models.Product) error {
	return config.DB.Create(p).Error
}

func (r *productRepo) Update(id int, p *models.Product) error {
	return config.DB.Model(&models.Product{}).Where("id = ?", id).Updates(p).Error
}

func (r *productRepo) Delete(id int) error {
	return config.DB.Delete(&models.Product{}, id).Error
}
