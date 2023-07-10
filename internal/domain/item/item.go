package item

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty"`
	UserID      string                 `bson:"user_id"` // Index field
	NormalItem  map[string]interface{} `bson:"normal_item,omitempty"`
	PremiumItem map[string]interface{} `bson:"premium_item,omitempty"`
	CreatedAt   time.Time              `bson:"created_at"`
	UpdatedAt   time.Time              `bson:"updated_at"`
}

func (i *Item) UpdateNormalItem(newItem map[string]interface{}) {
	i.NormalItem = newItem
}

func (i *Item) UpdatePremiumItem(newItem map[string]interface{}) {
	i.PremiumItem = newItem
}
