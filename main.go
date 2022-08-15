package main

import (
	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/wil-g2/unico-challenge/api"
	"github.com/wil-g2/unico-challenge/domain/fair"
	"github.com/wil-g2/unico-challenge/infra/config"
	"github.com/wil-g2/unico-challenge/infra/database"
	"github.com/wil-g2/unico-challenge/infra/log"
	"github.com/wil-g2/unico-challenge/infra/repositories"
	"github.com/wil-g2/unico-challenge/infra/validator"
)

func main() {

	envConfig := config.SetupEnvFile()

	mysql := database.InitMysql(envConfig)
	validator := validator.NewValidator(validatorv10.New())
	repo := repositories.NewFairRepository(mysql)
	service := fair.NewService(repo, validator)
	fairHandler := api.NewUserHandler(service)
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
	})

	log.Info("Fiber Server created with sucess", nil)
	api.SetupRoutes(app, fairHandler)
	if err := app.Listen(":" + envConfig.Port); err != nil {
		log.Error("Error listening", err, &log.LogContext{"port": envConfig.Port})

	}

}
