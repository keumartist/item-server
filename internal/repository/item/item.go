package item

import (
	"context"
	"time"

	"art-item/internal/domain/item"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBItemRepository struct {
	collection *mongo.Collection
}

func NewMongoDBItemRepository(db *mongo.Database, collectionName string) ItemRepository {
	return &MongoDBItemRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoDBItemRepository) FindByID(id string) (*item.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var result item.Item
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (r *MongoDBItemRepository) FindByUserID(userID string) (*item.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var result item.Item
	err := r.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (r *MongoDBItemRepository) UpdateNormalItem(id string, normalItem map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"normal_item": normalItem}})
	return err
}

func (r *MongoDBItemRepository) UpdatePremiumItem(id string, premiumItem map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"premium_item": premiumItem}})
	return err
}
