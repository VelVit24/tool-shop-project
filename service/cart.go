package service

import (
	"errors"

	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/repository"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) CreateCart(id_user int, cart *models.Cart) error {
	stock, err := s.repo.SelectProductStock(cart.Id_product)
	if err != nil {
		return err
	}
	if cart.Amount > stock {
		return errors.New("Not enough stock")
	}
	err = s.repo.InsertCart(id_user, cart)
	return err
}
func (s *CartService) UpdateCart(id_user int, cart *models.Cart) error {
	stock, err := s.repo.SelectProductStock(cart.Id_product)
	if err != nil {
		return err
	}
	if cart.Amount > stock {
		return errors.New("Not enough stock")
	}
	err = s.repo.UpdateCart(id_user, cart)
	return err
}
func (s *CartService) DeleteCart(id_user int, id int) error {
	err := s.repo.DeleteCart(id_user, id)
	return err
}
func (s *CartService) GetCart(id_user int) ([]models.CartItems, error) {
	carts, err := s.repo.SelectCart(id_user)
	return carts, err
}
