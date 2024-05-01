package main

import (
	"github.com/labstack/echo/v4"
	"math/rand/v2"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, Go Workshop!"})
	})
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "pong", "version": version()})
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func version() string {
	return string(rand.Int32())
}
