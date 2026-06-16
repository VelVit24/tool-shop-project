package service

import (
	"strconv"

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
func (s *ProductService) GetProduct(page, limit string) ([]models.Product, error) {
	if page == "" || limit == "" {
		products, err := s.repo.SelectProduct()
		return products, err
	} else {
		p, err := strconv.Atoi(page)
		if err != nil {
			return nil, err
		}
		l, err := strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
		products, err := s.repo.SelectProductPagination(p, l)
		return products, err
	}

}
func (s *ProductService) GetProductId(id int) (models.Product, error) {
	product, err := s.repo.SelectProductId(id)
	return product, err
}
