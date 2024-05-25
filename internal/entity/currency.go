package entity

import "github.com/shopspring/decimal"

// в настоящем проекте я бы еще учел следующие вещи:
// - отделить entity от моделей которые принимает слой репозитория
// - добавить в репо модели created_at и updated_at
type Currency struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Code      string `gorm:"size:10;uniqueIndex"`
	Available bool
	Rate      decimal.Decimal `gorm:"type:decimal(20,5);"`
	Type      string
}

type ConversionResult struct {
	From   string
	To     string
	Amount float64
	Result float64
}
