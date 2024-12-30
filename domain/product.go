package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempity" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Price       int                `bson:"price" json:"price"`
	Category    string             `bson:"category" json:"category"`
	Discription string             `bson:"discription" json:"discription"`
	Imagepath   string             `bson:"image_path" json:"image_path"`
}

type ProductUseCase interface {
	CreateProduct(product Product) (Product, error)
	GetAllProduct() ([]Product, error)
	GetProductByID(id string) (Product, error)
	UpdateProduct(product Product, id string) (Product, error)
	DeleteProduct(id string) (Product, error)
}

type ProductRepository interface {
	CreateProduct(product Product) (Product, error)
	GetAllProduct() ([]Product, error)
	GetProductByID(id string) (Product, error)
	UpdateProduct(product Product, id string) (Product, error)
	DeleteProduct(id string) (Product, error)
}
