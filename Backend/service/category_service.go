package service

import (
	"github.com/luthfi/pos/models"
	"github.com/luthfi/pos/repository"
)

type CategoryService struct{ repo repository.CategoryRepository }

func NewCategoryService(r repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: r}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Create(c *models.Category) error {
	return s.repo.Create(c)
}

func (s *CategoryService) Update(id uint, c *models.Category) error {
	return s.repo.Update(id, c)
}

func (s *CategoryService) Delete(id uint) error {
	return s.repo.Delete(id)
}
