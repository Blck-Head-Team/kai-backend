package main

import (
	"kai-backend/app/middlewares"
	"kai-backend/app/routes"
	_newslettersUsecase "kai-backend/business/newsletters"
	_newslettersController "kai-backend/controllers/newsletters"
	"kai-backend/repository/mysql"
	_newslettersRepo "kai-backend/repository/mysql/newsletters"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if viper.GetBool("debug") {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	e := echo.New()
	db := mysql.InitDB()
	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString("JWT_SECRET"),
		ExpiresDuration: viper.GetInt("JWT_EXPIRED"),
	}

	timeoutContext := time.Duration(viper.GetInt("SERVER_TIMEOUT")) * time.Second

	newslettersRepo := _newslettersRepo.NewNewsletterRepository(db)
	newslettersUsecase := _newslettersUsecase.NewNewsletterUsecase(newslettersRepo, timeoutContext)
	newslettersController := _newslettersController.NewHandler(newslettersUsecase)

	routesInit := routes.ControllersList{
		JWTMiddleware:         configJWT.Init(),
		NewslettersController: newslettersController,
	}

	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(viper.GetString("SERVER_PORT")))
}
