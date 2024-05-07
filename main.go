package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Menggunakan fungsi Static untuk mengakses file statis seperti index.html
	e.Static("/", "public")

	e.Logger.Fatal(e.Start(":8081"))
}
