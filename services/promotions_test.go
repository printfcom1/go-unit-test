package services_test

import (
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

	//Arrage

}
