package golangecho

import (
	// "fmt"
	"net/http"
	// "os"
	// "strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var products = []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}
