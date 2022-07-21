# nrfiber

`nrfiber` is one of [New Relic Go Agent](https://github.com/newrelic/go-agent) integration packages.
It instruments inbound requests through the [Fiber](https://gofiber.io/) framework.

[![Go Reference](https://pkg.go.dev/badge/github.com/ichizero/nrfiber.svg)](https://pkg.go.dev/github.com/ichizero/nrfiber)
[![Go Report Card](https://goreportcard.com/badge/github.com/ichizero/nrfiber)](https://goreportcard.com/report/github.com/ichizero/nrfiber)

## 🚀 Install

```bash
go get -u github.com/ichizero/nrfiber
```

## 🧐 Usage

```go
package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ichizero/nrfiber"
	"github.com/newrelic/go-agent/v3/newrelic"
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
		return c.Status(http.StatusOK).Send(c.Request().Body())
	})

	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}
```
