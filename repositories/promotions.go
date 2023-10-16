package repositories

type PromptionRepository interface {
	GetPromotion() (Promotion, error)
}

type Promotion struct {
	ID              string
	PurchaseMin     int
	DiscountPercent int
}
