package http

import (
	"errors"

	customErrors "exchange-rate-calculator/internal/errors"

	"github.com/gofiber/fiber/v2"
)

func mapErrorToStatus(err error) (int, string) {
	// для ошибок можно добавить более кастомные настройки
	// например указывать какая именно валюта недоступна/не найдена
	switch {
	case errors.Is(err, customErrors.ErrCurrencyNotFound):
		return fiber.StatusNotFound, err.Error()
	case errors.Is(err, customErrors.ErrCurrencyNotAvailable):
		return fiber.StatusBadRequest, err.Error()
	case errors.Is(err, customErrors.ErrCurrencyCanNotBeCalculated):
		return fiber.StatusPreconditionFailed, err.Error()
	default:
		return fiber.StatusInternalServerError, "Internal server error"
	}
}
