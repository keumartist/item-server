package item

import (
	"context"
	"time"

	"art-item/internal/domain/item"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var result item.Item
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
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

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": bson.M{"normal_item": normalItem, "updated_at": time.Now()}})
	return err
}

func (r *MongoDBItemRepository) UpdatePremiumItem(id string, premiumItem map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": bson.M{"premium_item": premiumItem, "updated_at": time.Now()}})
	return err
}

func (r *MongoDBItemRepository) CreateItem(userID string, normalItem map[string]interface{}, premiumItem map[string]interface{}) (*item.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newItem := &item.Item{
		UserID:      userID,
		NormalItem:  normalItem,
		PremiumItem: premiumItem,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	res, err := r.collection.InsertOne(ctx, newItem)
	if err != nil {
		return nil, err
	}

	newItem.ID = res.InsertedID.(primitive.ObjectID)

	return newItem, nil
}
