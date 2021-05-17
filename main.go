package main

import (
	bookHandler "bibo.api/api/book"
	userHandler "bibo.api/api/user"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
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

	e.GET("/books", bookHandler.FindAll)
	e.POST("/book", bookHandler.CreateBook)
	e.GET("/book/:isbn", bookHandler.FindOne)
	e.PUT("/book/:isbn", bookHandler.Update)
	//e.DELETE("/book/:isbn", bookHandler.Delete)

	userRoute := e.Group("/user")
	userRoute.POST("/create", userHandler.CreateUser)
	userRoute.POST("/login", userHandler.AuthenticateUser)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
