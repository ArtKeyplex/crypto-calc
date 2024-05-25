package build

import (
	"exchange-rate-calculator/configs"
)

type Builder struct {
	config configs.Config

	shutdown shutdown
}

func New(conf configs.Config) *Builder {
	b := Builder{config: conf}

	return &b
}
