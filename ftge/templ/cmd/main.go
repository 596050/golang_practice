package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

app.GET("/user", handler.)
	app.Start(":8080")

	fmt.Println("Hello, World!")
}
