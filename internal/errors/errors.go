package errors

import "errors"

var (
	ErrCurrencyNotFound           = errors.New("currency not found")
	ErrCurrencyNotAvailable       = errors.New("currency not available for exchange")
	ErrCurrencyCanNotBeCalculated = errors.New("currency can not be calculated")
)
