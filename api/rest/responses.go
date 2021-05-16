package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ResponseSuccess (c echo.Context, payload interface{}) error {
	return c.JSON(http.StatusOK, payload)
}

func ResponseCreated (c echo.Context, payload interface{}) error {
	return c.JSON(http.StatusCreated, payload)
}

func ResponseBadRequest(c echo.Context, payload interface{}) error {
	return c.JSON(http.StatusBadRequest, payload)
}

func ResponseNotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "")
}

func ResponseInternalError(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, "")
}

func ResponseForbidden(c echo.Context) error {
	return c.JSON(http.StatusForbidden, "")
}