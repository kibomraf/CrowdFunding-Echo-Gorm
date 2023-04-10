package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"crowdfunding/auth"
	"crowdfunding/campaign"
	"crowdfunding/handler"
	"crowdfunding/users"
)

func main() {

	dsn := "host=localhost user=root password=123qwe dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//users
	userRepo := users.UserRepository(db)
	userService := users.UserService(userRepo)
	authService := auth.NewJWTservice("CrowdFunding-Echo")
	userHandler := handler.UserHandler(userService, authService)
	//campaign
	campRepo := campaign.CampaignRepository(db)
	campaign, err := campRepo.FindUserById(2)
	if err != nil {
		panic(err)
	}
	for _, c := range campaign {
		fmt.Println("Campaign: ", c.Name)
		fmt.Println("Campaign: ", c.CampaignImages[0].FileName)

	}
	app := echo.New()
	//midlerware
	jwtMiddleWare := authService.JWTMiddleware(authService, userService)
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	//route user
	api := app.Group("/api/v1")
	api.POST("/user", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.GET("/user/fetch", userHandler.FetchUser)
	api.POST("/user/check", userHandler.CheckEmailAvailablity)
	api.POST("/user/avatar", userHandler.UploadAvatar, jwtMiddleWare)

	app.Logger.Fatal(app.Start(":8080"))
}
