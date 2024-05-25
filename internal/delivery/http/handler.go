package http

import (
	"fmt"

	"exchange-rate-calculator/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Handler struct {
	usecase usecase.CurrencyUsecase
	log     *zerolog.Logger
}

func NewHandler(app *fiber.App, usecase usecase.CurrencyUsecase, log *zerolog.Logger) {
	handler := &Handler{
		usecase: usecase,
		log:     log,
	}

	app.Post("/convert", handler.ConvertCurrency)
}

type ConvertCurrencyRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

// ErrorResponse представляет ответ с ошибкой.
type ErrorResponse struct {
	Error string `json:"error"`
}

// ConvertCurrency обрабатывает запрос на конвертацию валют
// @Summary Convert Currency
// @Description Convert currency from one to another
// @Tags currency
// @Accept json
// @Produce json
// @Param request body ConvertCurrencyRequest true "Request body"
// @Success 200 {object} entity.ConversionResult
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /convert [post].
func (h *Handler) ConvertCurrency(c *fiber.Ctx) error {
	var req ConvertCurrencyRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result, err := h.usecase.ConvertCurrency(c.Context(), req.From, req.To, req.Amount)
	if err != nil {
		status, message := mapErrorToStatus(err)

		return c.Status(status).JSON(fiber.Map{
			"error": message,
		})
	}

	err = c.JSON(result)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %w", err)
	}

	return nil
}
