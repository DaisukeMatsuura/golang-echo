package app

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var e = echo.New()
var v = validator.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configration")
	}
}

// Starts the application
func Start() {
	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.PUT("/products/:id", updateProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port)))
}
