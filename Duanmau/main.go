package main

import (
	"duanmau/database"
	"duanmau/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	database.InitDB()
	defer database.DB.Close()

	userHandler := handler.UserHandler{
		UserRepo : repo_impl.NewUserRepo(sql),
	}
api:= router.API{
	Echo: e,
	UserHandler: userHandler,
}
api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))

}

