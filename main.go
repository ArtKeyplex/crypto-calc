package main

import (
	"context"
	"exchange-rate-calculator/cmd"
	"exchange-rate-calculator/configs"
	"exchange-rate-calculator/pkg/logger"
	"os"

	_ "exchange-rate-calculator/docs/swagger"
)

func main() {
	conf, err := configs.LoadConfig()
	if err != nil {
		// в тз сказано никаких паник, но думаю, при изначальной конфигурации приложения, если критически важные данные
		// не заданы, то можно паникнуть. Например, подключение к бд, прочтение енвов и тд
		panic(err)
	}

	log := logger.New(conf.LogLevel)
	ctx := log.WithContext(context.Background())

	log.Info().Msg("The application is launching")

	exitCode := 0

	err = cmd.Run(ctx, conf)
	if err != nil {
		log.Err(err).Send()

		exitCode = 1
	}

	os.Exit(exitCode)
}
