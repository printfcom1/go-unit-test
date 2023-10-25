// build: intrgration
package handler_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-unit-test/handler"
	"github.com/go-unit-test/repositories"
	"github.com/go-unit-test/services"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscountIntegration(t *testing.T) {
	t.Run("susscess", func(t *testing.T) {
		amount := 100
		expect := 80

		promoRepo := repositories.NewPromotionRipositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              "1",
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expect, nil)

		promoHandler := handler.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)
		defer res.Body.Close()

		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expect), string(body))
		}

	})
}
