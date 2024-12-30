package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	ID       primitive.ObjectID `bson:"_id,omitempity" json:"id"`
	UserID   string             `bson:"user_id" json:"user_id"`
	Products []Product
}
type CartRepository interface {
	CreateCart(cart Cart) (Cart, error)
	GetCartByUserID(userID string) (Cart, error)
	UpdateCart(cart Cart) (Cart, error)
	DeleteCart(userID string) error
}

type CartUseCase interface {
	CreateCart(cart Cart) (Cart, error)
	GetCartByUserID(userID string) (Cart, error)
	UpdateCart(cart Cart) (Cart, error)
	DeleteCart(userID string) error
}
