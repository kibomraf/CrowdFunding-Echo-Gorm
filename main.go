package main

import (
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
	//auth
	authService := auth.NewJWTservice("CrowdFunding-Echo")

	//users
	userRepo := users.UserRepository(db)
	userService := users.UserService(userRepo)
	userHandler := handler.UserHandler(userService, authService)

	//campaign
	campRepo := campaign.CampaignRepository(db)
	campService := campaign.Campaignservices(campRepo)
	campHandler := handler.NewCampHandler(campService)

	app := echo.New()
	app.Static("/avatar", "./images/avatar/user/")

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

	//route camp
	api.GET("/campaigns", campHandler.GetCampaignc)
	api.GET("/campaigns/:id", campHandler.GetDetailsCampaign)
	api.POST("/campaigns", campHandler.CreateCampaign, jwtMiddleWare)
	api.PUT("/campaigns/:id", campHandler.UpdateCampaign, jwtMiddleWare)
	api.POST("/campaigns/image", campHandler.UploadImages, jwtMiddleWare)

	app.Logger.Fatal(app.Start(":8080"))
}
