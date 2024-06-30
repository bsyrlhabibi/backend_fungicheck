package user

import (
	"fastfooducate/module/feature/user/domain"
	"fastfooducate/module/feature/user/handler"
	"fastfooducate/module/feature/user/repository"
	"fastfooducate/module/feature/user/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	repo domain.UserRepositoryInterface
	serv domain.UserServiceInterface
	hand domain.UserHandlerInterface
)

func InitializeUser(db *gorm.DB) {

	repo = repository.NewUserRepository(db)
	serv = service.NewUserService(repo)
	hand = handler.NewUserHandler(serv)
}

func SetupRoutesUser(app *fiber.App) {
	api := app.Group("/api/v1/user")
	api.Get("/:id", hand.GetUserByID)
}
