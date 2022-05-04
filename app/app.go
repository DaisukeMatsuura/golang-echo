package app

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside custom middleware")
		c.Request().URL.Path = "/"
		fmt.Printf("%+v\n", c.Request())
		return next(c)
	}
}

// Starts the application
func Start() {
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.POST("product", createProduct, middleware.BodyLimit("1K"))
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port)))
}
