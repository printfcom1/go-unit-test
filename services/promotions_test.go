package services_test

import (
	"errors"
	"testing"

	"github.com/go-unit-test/repositories"
	"github.com/go-unit-test/services"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	type testCase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	cases := []testCase{
		{name: "applid 100", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "applid 200", purchaseMin: 100, discountPercent: 20, amount: 200, expected: 160},
		{name: "applid 300", purchaseMin: 100, discountPercent: 20, amount: 300, expected: 240},
		{name: "no applid 50", purchaseMin: 100, discountPercent: 20, amount: 50, expected: 50},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			promoRepo := repositories.NewPromotionRipositoryMock()
			promoRepo.On("GetPromotion").Return(repositories.Promotion{ID: "1", PurchaseMin: c.purchaseMin, DiscountPercent: c.discountPercent}, nil)
			promoService := services.NewPromotionService(promoRepo)

			calculate, _ := promoService.CalculateDiscount(c.amount)
			expected := c.expected

			assert.Equal(t, expected, calculate)
		})
	}

	t.Run("purchase amount zero", func(t *testing.T) {
		promoRepo := repositories.NewPromotionRipositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{ID: "1", PurchaseMin: 100, DiscountPercent: 20}, nil)
		promoService := services.NewPromotionService(promoRepo)

		_, err := promoService.CalculateDiscount(0)

		assert.ErrorIs(t, err, services.ErrZeroAmount)
		promoRepo.AssertNotCalled(t, "GetPromotion")
	})

	t.Run("repository error", func(t *testing.T) {
		promoRepo := repositories.NewPromotionRipositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, errors.New("error"))
		promoService := services.NewPromotionService(promoRepo)

		_, err := promoService.CalculateDiscount(100)

		assert.ErrorIs(t, err, services.ErrRepository)
	})
}
