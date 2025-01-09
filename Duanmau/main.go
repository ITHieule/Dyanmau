package main

import (
	"duanmau/database"
	"duanmau/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	database.InitDB()
	defer database.DB.Close()

	e := echo.New()
	e.GET("/", handler.Wellcom)
	e.GET("/user/login", handler.Dangnhap)
	e.GET("/user/login-sign-up", handler.Dangki)
	e.Logger.Fatal(e.Start(":3000"))
}
