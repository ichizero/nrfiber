package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/ichizero/nrfiber"
)

func main() {
	nrApp, err := newrelic.NewApplication(
		newrelic.ConfigEnabled(true),
		newrelic.ConfigAppName("nrfiber-example"),
		newrelic.ConfigLicense("license-key"),
	)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(nrfiber.Middleware(nrApp))

	app.Post("/echo", func(c *fiber.Ctx) error {
		sleepWithFiberContext(c)
		sleepWithContext(c.UserContext())
		return c.Status(http.StatusOK).Send(c.Request().Body())
	})

	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}

func sleepWithFiberContext(c *fiber.Ctx) {
	seg := nrfiber.FromContext(c).StartSegment("sleep-with-fiber-context")
	defer seg.End()
	time.Sleep(500 * time.Millisecond)
}

func sleepWithContext(ctx context.Context) {
	seg := newrelic.FromContext(ctx).StartSegment("sleep-with-context")
	defer seg.End()
	time.Sleep(500 * time.Millisecond)
}
