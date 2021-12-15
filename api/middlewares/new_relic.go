package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	. "github.com/newrelic/go-agent"
	"log"
)

const (
	NewrelicTxn = "newrelic-txn"
)

func SetupNewRelic(appName string, licenseKey string) echo.MiddlewareFunc {
	app, err := NewApplication(NewConfig(appName, licenseKey))

	if err != nil {
		log.Fatalf("New Relic: %s", err)
	}

	return BuildNewRelicApp(app)
}

func BuildNewRelicApp(app Application) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			transactionName := fmt.Sprintf("%s [%s]", c.Path(), c.Request().Method)
			txn := app.StartTransaction(transactionName, c.Response().Writer, c.Request())
			defer txn.End()

			c.Set(NewrelicTxn, txn)

			err := next(c)

			if err != nil {
				_ = txn.NoticeError(err)
			}

			return err
		}
	}
}
