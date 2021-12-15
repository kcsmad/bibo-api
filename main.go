package main

import (
	mw "bibo.api/api/middlewares"
	"bibo.api/api/v1/books"
	"bibo.api/api/v1/db"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"os"
)

var dbInstance = db.Repository{}
var dbURL = ""

var mainLogger = log.WithFields(log.Fields{
	"[domain]": "[Main]",
	"[env]":    "[" + os.Getenv("APP_ENV") + "]",
})

func main() {
	setupEnv()

	dbInstance.Connect(dbURL)
	defer dbInstance.Disconnect()

	echoInstance := setupEcho()
	v1Route := echoInstance.Group("/v1")
	setupBookHandler(v1Route)

	serve(echoInstance)
}

func setupBookHandler(e *echo.Group) {
	bookDAO := books.DAO{
		Repo: &dbInstance,
	}

	bookHandler := books.Handler{
		DAO: bookDAO,
	}

	route := e.Group("/books")
	route.GET("", bookHandler.FindAll)
	route.GET("/", bookHandler.FindAll)
}

func setupEnv() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()

		if err != nil {
			mainLogger.WithFields(log.Fields{
				"[method]": "[SetupEnv]",
				"[action]": "[Load Env File]",
			}).Fatal(err.Error())
		}

		dbURL = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") +
			"@tcp(" + os.Getenv("DB_HOST") + "" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")
	} else {
		dbURL = os.Getenv("CLEARDB_URL")
	}
}

func setupEcho() *echo.Echo {
	e := echo.New()
	e.Use(emw.CORS())
	e.Use(emw.Recover())
	e.Use(mw.SetupNewRelic(os.Getenv("NR_APP_NAME"), os.Getenv("NR_LICENSE_KEY")))
	e.Use(mw.Logrus())

	return e
}

func serve(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
