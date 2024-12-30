package controller

import (
	"fmt"
	"onlineshopbackend/domain"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase domain.ProductUseCase
}

func NewProductController(ProductUsecase domain.ProductUseCase) *ProductController {
	return &ProductController{
		ProductUsecase: ProductUsecase,
	}
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid request", "error": err.Error()})
		return
	}
	createdProduct, err := p.ProductUsecase.CreateProduct(product)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to create product", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Product created successfully", "data": createdProduct})
}

func (p *ProductController) GetAllProduct(c *gin.Context) {
	products, err := p.ProductUsecase.GetAllProduct()
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to retrieve products", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Products retrieved successfully", "data": products})
}

func (p *ProductController) GetProductByID(c *gin.Context) {
	id := c.Query("ID")
	product, err := p.ProductUsecase.GetProductByID(id)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to retrieve product", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Product retrieved successfully", "data": product})
}

func (p *ProductController) UpdateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid request", "error": err.Error()})
		return
	}
	id := c.Query("ID")

	product, err := p.ProductUsecase.UpdateProduct(product, id)
	fmt.Println(product.ID, id)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to update product", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Product updated successfully", "data": product})
}

func (p *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Query("ID")
	deletedProduct, err := p.ProductUsecase.DeleteProduct(id)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to delete product", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Product deleted successfully", "data": deletedProduct})
}
