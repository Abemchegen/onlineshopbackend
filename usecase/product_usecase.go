package usecase

import (
	"errors"
	"onlineshopbackend/domain"
)

type productUseCase struct {
	productRepository domain.ProductRepository
}

func NewProductUseCase(productRepository domain.ProductRepository) domain.ProductUseCase {
	return &productUseCase{
		productRepository: productRepository,
	}
}

func (p *productUseCase) CreateProduct(product domain.Product) (domain.Product, error) {
	if product.Name == "" {
		return domain.Product{}, errors.New("Name is required")
	}
	if product.Price == 0 {
		return domain.Product{}, errors.New("BuyPrice is required")
	}

	_, err := p.productRepository.CreateProduct(product)
	return product, err
}

func (p *productUseCase) GetAllProduct() ([]domain.Product, error) {
	return p.productRepository.GetAllProduct()
}

func (p *productUseCase) GetProductByID(id string) (domain.Product, error) {
	return p.productRepository.GetProductByID(id)
}

func (p *productUseCase) UpdateProduct(product domain.Product, id string) (domain.Product, error) {
	if product.Name == "" {
		return domain.Product{}, errors.New("name is required")
	}
	if product.Price == 0 {
		return domain.Product{}, errors.New("price is required")
	}

	product, err := p.productRepository.UpdateProduct(product, id)
	return product, err
}

func (p *productUseCase) DeleteProduct(id string) (domain.Product, error) {
	product, err := p.productRepository.GetProductByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	_, err = p.productRepository.DeleteProduct(id)
	return product, err
}
