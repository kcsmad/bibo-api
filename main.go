package main

import (
	userHandler "bibo.api/api/user"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		setupEnv()
	}

	startup()
}

func setupEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading env file.")
	}
}

func startup()  {
	// TODO: New Relic Configuration
	serve()
}

func serve() {
	e := echo.New()

	bookRoute := e.Group("/books")

	bookRoute.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == os.Getenv("AUTH_SECRET"), nil
	}))

	userRoute := e.Group("/user")
	userRoute.POST("/create", userHandler.CreateUser)
	userRoute.POST("/login", userHandler.AuthenticateUser)

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Kek")
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
