package router

 type API struct {

	Echo * echo.Echo
	UserHandler handler.UserHandler
 }
 func SetupRouter(){
	api.Echo.GET("/api/users",api.UserHandler.User )	
	api.Echo.GET("/user/login",api.UserHandler.Dangnhap)
	api.Echo.GET("/user/login-sign-up", api.UserHandler.Dangki)
 }