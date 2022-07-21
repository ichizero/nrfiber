# nrfiber

## Install

```shell
go get -u github.com/ichizero/nrfiber
```

## Usage

```go
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
