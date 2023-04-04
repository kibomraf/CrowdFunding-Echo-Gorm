package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"crowdfunding/auth"
	"crowdfunding/handler"
	"crowdfunding/users"
)

func main() {

	dsn := "host=localhost user=root password=123qwe dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	userRepo := users.UserRepository(db)
	userService := users.UserService(userRepo)
	userHandler := handler.UserHandler(userService)
	authService := auth.AuthService()
	tkn, _ := authService.GenerateToken(1)
	fmt.Println("Token :", tkn)

	app := echo.New()
	//midlerware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	//route user
	api := app.Group("/api/v1")
	api.POST("/user", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.GET("/user/fetch", userHandler.FetchUser)
	api.POST("/user/check", userHandler.CheckEmailAvailablity)
	api.POST("/user/avatar", userHandler.UploadAvatar)

	app.Logger.Fatal(app.Start(":8080"))
}
