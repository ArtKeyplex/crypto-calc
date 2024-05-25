package build

import (
	"context"

	_ "exchange-rate-calculator/docs/swagger"
	"exchange-rate-calculator/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog"
)

// @title Exchange Rate Calculator API
// @version 1.0
// @description API для калькулятора обменного курса между крипто и фиатными валютами
// @host localhost:8080
// @BasePath /.
func (b *Builder) RestAPIServer(ctx context.Context) (*fiber.App, error) {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	currencyUsecase, err := b.CurrencyUsecase()
	if err != nil {
		return nil, err
	}

	http.NewHandler(app, currencyUsecase, zerolog.Ctx(ctx))

	return app, nil
}
