package service

import (
	"github.com/VelVit24/projext/dto"
	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(id_user int, order *models.Order) error {
	err := s.repo.InsertOrder(id_user, order)
	return err
}

func (s *OrderService) UpdateOrder(order *models.Order) error {
	err := s.repo.UpdateOrder(order)
	return err
}

func (s *OrderService) DeleteOrder(id int) error {
	err := s.repo.DeleteOrder(id)
	return err
}

func (s *OrderService) GetOrders(id_user int, role string, page, limit string) ([]dto.OrderView, error) {
	p, l, err := PaginationParse(page, limit)
	if err != nil {
		return nil, err
	}
	orders, err := s.repo.SelectOrders(id_user, p, l, role)
	return orders, err
}
