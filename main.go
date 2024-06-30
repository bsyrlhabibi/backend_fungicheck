package main

import (
	"fastfooducate/config"
	"fastfooducate/module/feature/middleware"
	"fastfooducate/module/feature/route"
	"fastfooducate/module/feature/user/repository"
	"fastfooducate/module/feature/user/service"
	"fastfooducate/utils/database"
	"fastfooducate/utils/token"
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
