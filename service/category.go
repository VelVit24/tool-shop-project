package service

import (
	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	err := s.repo.InsertCategory(category)
	return err
}
func (s *CategoryService) UpdateCategory(category *models.Category) error {
	err := s.repo.UpdateCategory(category)
	return err
}
func (s *CategoryService) DeleteCategory(id int) error {
	err := s.repo.DeleteCategory(id)
	return err
}
func (s *CategoryService) GetCategory() ([]models.Category, error) {
	categories, err := s.repo.SelectCategories()
	return categories, err
}
