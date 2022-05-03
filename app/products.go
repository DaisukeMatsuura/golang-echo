package app

import (
	// "fmt"
	"net/http"
	"strconv"

	// "os"

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

func getProduct(c echo.Context) error {
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for _, p := range products {
		for k := range p {
			if pID == k {
				product = p
			}
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found.")
	}
	return c.JSON(http.StatusOK, product)
}

func updateProduct(c echo.Context) error {
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for _, p := range products {
		for k := range p {
			if pID == k {
				product = p
			}
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found.")
	}

	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body
	e.Validator = &ProductValidator{validator: v}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}

	product[pID] = reqBody.Name
	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	var product map[int]string
	var index int
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil
	}
	for i, p := range products {
		for k := range p {
			if pID == k {
				product = p
				index = i
			}
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found.")
	}
	splice := func(s []map[int]string, index int) []map[int]string {
		return append(s[:index], s[index+1:]...)
	}
	products = splice(products, index)

	return c.JSON(http.StatusOK, product)
}
