package service

import (
	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	err := s.repo.InsertProduct(product)
	return err
}
func (s *ProductService) UpdateProduct(product *models.Product) error {
	err := s.repo.UpdateProduct(product)
	return err
}
func (s *ProductService) DeleteProduct(id int) error {
	err := s.repo.DeleteProduct(id)
	return err
}
func (s *ProductService) GetProduct() ([]models.Product, error) {
	products, err := s.repo.SelectProduct()
	return products, err
}
