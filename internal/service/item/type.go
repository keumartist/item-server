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
