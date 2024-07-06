package main

import (
	"fungicheck/config"
	"fungicheck/module/feature/middleware"
	"fungicheck/module/feature/route"
	"fungicheck/module/feature/user/repository"
	"fungicheck/module/feature/user/service"
	"fungicheck/utils/database"
	"fungicheck/utils/token"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	var initConfig = config.InitConfig()
	db := database.InitPGSDatabase(*initConfig)
	jwtService := token.NewJWT(initConfig.Secret)

	middleware.SetupMiddlewares(app)
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	database.Migrate(db)
	route.SetupRoutes(app, db, jwtService, userService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, FungiCheck!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	err := app.Listen(":" + port)
	if err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
