package nrfiber

import (
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Middleware creates Fiber middleware that instruments requests.
//
// 	app := fiber.New()
// 	// Add the nrfiber middleware before other middlewares or routes:
// 	app.Use(nrfiber.Middleware(app))
func Middleware(app *newrelic.Application, opts ...Option) fiber.Handler {
	if app == nil {
		return func(c *fiber.Ctx) error { return c.Next() } // no-op
	}

	cfg := newConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	return func(c *fiber.Ctx) error {
		txn := app.StartTransaction(cfg.transactionNameFormatter(c))
		defer txn.End()

		txn.SetWebRequestHTTP(convertToRequest(c))
		c.SetUserContext(newrelic.NewContext(c.UserContext(), txn))

		err := c.Next()
		if err != nil {
			txn.NoticeError(err)
		}
		txn.SetWebResponse(nil).WriteHeader(c.Response().StatusCode())

		return nil
	}
}

func convertToHeader(c *fiber.Ctx) *http.Header {
	rc := c.Context()
	h := http.Header{}
	rc.Request.Header.VisitAll(func(k, v []byte) {
		h.Set(string(k), string(v))
	})
	return &h
}

func convertToRequest(c *fiber.Ctx) *http.Request {
	rc := c.Context()
	rURL, _ := url.ParseRequestURI(string(rc.RequestURI()))
	return &http.Request{
		Header: *convertToHeader(c),
		URL:    rURL,
		Method: string(rc.Method()),
		TLS:    rc.TLSConnectionState(),
		Host:   string(rc.Host()),
	}
}

// FromContext returns *newrelic.Transaction from *fiber.Ctx if present.
func FromContext(c *fiber.Ctx) *newrelic.Transaction {
	return newrelic.FromContext(c.UserContext())
}
