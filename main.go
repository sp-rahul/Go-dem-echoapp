package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Echo")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello aWorld. I am now in ECHO")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
