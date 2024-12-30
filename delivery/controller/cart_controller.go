package controller

import (
	"onlineshopbackend/domain"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	CartUsecase domain.CartUseCase
}

func NewCartController(CartUsecase domain.CartUseCase) *CartController {
	return &CartController{
		CartUsecase: CartUsecase,
	}
}

func (p *CartController) CreateCart(c *gin.Context) {
	var Cart domain.Cart
	if err := c.ShouldBindJSON(&Cart); err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid request", "error": err.Error()})
		return
	}
	Cart.UserID = c.Query("user_id")
	createdCart, err := p.CartUsecase.CreateCart(Cart)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to create Cart", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Cart created successfully", "data": createdCart})
}

func (p *CartController) GetCartByUserID(c *gin.Context) {
	id := c.Query("id")

	Cart, err := p.CartUsecase.GetCartByUserID(id)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to retrieve Cart", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Cart retrieved successfully", "data": Cart})
}

func (p *CartController) UpdateCart(c *gin.Context) {
	var Cart domain.Cart
	if err := c.ShouldBindJSON(&Cart); err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid request", "error": err.Error()})
		return
	}
	Cart.UserID = c.GetString("user_id")
	updatedCart, err := p.CartUsecase.UpdateCart(Cart)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to update Cart", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Cart updated successfully", "data": updatedCart})
}

func (p *CartController) DeleteCart(c *gin.Context) {
	id := c.Query("id")
	err := p.CartUsecase.DeleteCart(id)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to delete Cart", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Cart deleted successfully"})
}
