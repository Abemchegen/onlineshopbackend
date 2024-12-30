package usecase

import (
	"onlineshopbackend/domain"
)

type CartUseCase struct {
	CartRepository domain.CartRepository
}

func NewCartUseCase(CartRepository domain.CartRepository) domain.CartUseCase {
	return &CartUseCase{
		CartRepository: CartRepository,
	}
}

func (c *CartUseCase) CreateCart(Cart domain.Cart) (domain.Cart, error) {
	_, err := c.CartRepository.CreateCart(Cart)
	return Cart, err
}

func (c *CartUseCase) GetCartByUserID(id string) (domain.Cart, error) {
	var cart domain.Cart
	cart, err := c.CartRepository.GetCartByUserID(id)
	return cart, err
}

func (c *CartUseCase) UpdateCart(Cart domain.Cart) (domain.Cart, error) {

	cart, err := c.CartRepository.UpdateCart(Cart)
	return cart, err
}

func (c *CartUseCase) DeleteCart(id string) error {

	err := c.CartRepository.DeleteCart(id)
	return err
}
