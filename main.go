package main

import (
	// "fmt"
	// "net/http"
	// "os"
	// "strconv"

	// "github.com/go-playground/validator"
	// "github.com/labstack/echo/v4"
	"golang-echo/app"
)

// type ProductValidator struct {
// 	validator *validator.Validate
// }

// func (p *ProductValidator) Validate(i interface{}) error {
// 	return p.validator.Struct(i)
// }

func main() {
	app.Start()

	// e.POST("/products", func(c echo.Context) error {
	// 	type body struct {
	// 		Name string `json:"product_name" validate:"required,min=4"`
	// 	}
	// 	var reqBody body
	// 	e.Validator = &ProductValidator{validator: v}
	// 	if err := c.Bind(&reqBody); err != nil {
	// 		return err
	// 	}
	// 	if err := v.Struct(reqBody); err != nil {
	// 		return err
	// 	}

	// 	product := map[int]string{
	// 		len(products) + 1: reqBody.Name,
	// 	}
	// 	products = append(products, product)
	// 	return c.JSON(http.StatusOK, product)
	// })
}
