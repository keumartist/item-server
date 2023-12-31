package item

type GetItemByIDInput struct {
	ID string
}

type GetItemByUserIDInput struct {
	UserID string
}

type UpdateNormalItemInput struct {
	UserID  string
	NewItem map[string]interface{}
}

type UpdatePremiumItemInput struct {
	UserID  string
	NewItem map[string]interface{}
}

type CreateItemInput struct {
	UserID      string                 `json:"user_id"`
	NormalItem  map[string]interface{} `json:"normal_item"`
	PremiumItem map[string]interface{} `json:"premium_item"`
}
