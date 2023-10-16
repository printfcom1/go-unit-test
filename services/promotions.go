package services

import "github.com/go-unit-test/repositories"

type PromotionService interface {
	CalculateDiscount(int) (int, error)
}

type promotionService struct {
	promoRepo repositories.PromptionRepository
}

func NewPromotionService(promoRepo repositories.PromptionRepository) PromotionService {
	return promotionService{promoRepo: promoRepo}
}

func (s promotionService) CalculateDiscount(amount int) (int, error) {
	if amount <= 0 {
		return 0, ErrZeroAmount
	}

	promo, err := s.promoRepo.GetPromotion()
	if err != nil {
		return 0, ErrRepository
	}

	if amount >= promo.PurchaseMin {
		return amount - (promo.DiscountPercent * amount / 100), nil
	}

	return amount, nil
}
