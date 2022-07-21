package nrfiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type config struct {
	transactionNameFormatter func(c *fiber.Ctx) string
}

func newConfig() *config {
	return &config{
		transactionNameFormatter: func(c *fiber.Ctx) string {
			return fmt.Sprintf("%s %s", c.Request().Header.Method(), c.Path())
		},
	}
}

// Option is an option for Middleware.
type Option func(cfg *config)

// WithTransactionNameFormatter takes a function that determine a transaction name for each request.
func WithTransactionNameFormatter(f func(c *fiber.Ctx) string) Option {
	return func(cfg *config) {
		cfg.transactionNameFormatter = f
	}
}
