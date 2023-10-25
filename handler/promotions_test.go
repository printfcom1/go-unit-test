package handler_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-unit-test/handler"
	"github.com/go-unit-test/services"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("susscess", func(t *testing.T) {
		amount := 100
		expect := 80

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

	t.Run("status BadRequest", func(t *testing.T) {
		amount := 100
		expect := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expect, nil)

		promoHandler := handler.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", "a"), nil)
		res, _ := app.Test(req)
		defer res.Body.Close()

		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	t.Run("status StatusNotFound", func(t *testing.T) {
		amount := 0

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(0, services.ErrZeroAmount)

		promoHandler := handler.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)
		defer res.Body.Close()

		assert.Equal(t, fiber.StatusNotFound, res.StatusCode)
	})
}
