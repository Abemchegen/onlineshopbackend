package repo

import (
	"context"
	"onlineshopbackend/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	database   mongo.Database
	collection string
}

func NewProductRepository(database mongo.Database, collection string) domain.ProductRepository {
	return &ProductRepository{
		database:   database,
		collection: collection}

}

func (p *ProductRepository) CreateProduct(product domain.Product) (domain.Product, error) {
	objID := primitive.NewObjectID()
	product.ID = objID
	_, err := p.database.Collection(p.collection).InsertOne(context.Background(), product)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (p *ProductRepository) GetAllProduct() ([]domain.Product, error) {
	var products []domain.Product
	cursor, err := p.database.Collection(p.collection).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) GetProductByID(id string) (domain.Product, error) {
	var product domain.Product
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Product{}, err
	}
	err = p.database.Collection(p.collection).FindOne(context.Background(), bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (p *ProductRepository) UpdateProduct(product domain.Product, id string) (domain.Product, error) {
	product.ID, _ = primitive.ObjectIDFromHex(id)
	_, err := p.database.Collection(p.collection).UpdateOne(context.Background(), bson.M{"_id": product.ID}, bson.M{"$set": product})
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (p *ProductRepository) DeleteProduct(id string) (domain.Product, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Product{}, err
	}
	var product domain.Product
	err = p.database.Collection(p.collection).FindOneAndDelete(context.Background(), bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
