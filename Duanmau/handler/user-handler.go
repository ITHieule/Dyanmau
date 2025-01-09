package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Dangnhap(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "lan",
		"email": "lan@gmail.com",
	})
}
func Dangki(c echo.Context) error {

	type User struct {
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
		Age      int    `json:"age"`
	}
	user := User{
		Email:    "lan@gmail.com",
		Fullname: "Lan",
		Age:      20,
	}
	return c.JSON(http.StatusOK, user)
}
