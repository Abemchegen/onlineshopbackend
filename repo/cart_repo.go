package repo

import (
	"context"
	"onlineshopbackend/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartRepository struct {
	database   mongo.Database
	collection string
}

func NewCartRepository(database mongo.Database, collection string) domain.CartRepository {
	return &CartRepository{
		database:   database,
		collection: collection,
	}

}

func (c *CartRepository) CreateCart(Cart domain.Cart) (domain.Cart, error) {
	objID := primitive.NewObjectID()
	Cart.ID = objID
	_, err := c.database.Collection(c.collection).InsertOne(context.Background(), Cart)
	if err != nil {
		return domain.Cart{}, err
	}

	return Cart, nil
}

func (c *CartRepository) GetCartByUserID(id string) (domain.Cart, error) {
	var Cart domain.Cart
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Cart{}, err
	}
	err = c.database.Collection(c.collection).FindOne(context.Background(), bson.M{"_id": objID}).Decode(&Cart)
	if err != nil {
		return domain.Cart{}, err
	}
	return Cart, nil
}

func (c *CartRepository) UpdateCart(Cart domain.Cart) (domain.Cart, error) {
	_, err := c.database.Collection(c.collection).UpdateOne(context.Background(), bson.M{"_id": Cart.ID}, bson.M{"$set": Cart})
	if err != nil {
		return domain.Cart{}, err
	}
	return Cart, nil
}

func (c *CartRepository) DeleteCart(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	var Cart domain.Cart
	err = c.database.Collection(c.collection).FindOneAndDelete(context.Background(), bson.M{"_id": objID}).Decode(&Cart)
	if err != nil {
		return err
	}
	return nil
}
