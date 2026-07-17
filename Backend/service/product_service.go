package service

import (
	"github.com/luthfi/pos/models"
	"github.com/luthfi/pos/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) CreateProduct(p *models.Product) error {

	if p.Category != "Makanan" && p.Category != "Minuman" {
		p.ExpiryDate = nil
	}
	return s.repo.Create(p)
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) UpdateProduct(id int, p *models.Product) error {
	if p.Category != "Makanan" && p.Category != "Minuman" {
		p.ExpiryDate = nil
	}
	return s.repo.Update(id, p)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *ProductService) Search(query string) ([]models.Product, error) {
	return s.repo.Search(query)
}
