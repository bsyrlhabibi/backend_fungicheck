package auth

import (
	"fastfooducate/module/feature/auth/domain"
	utils "fastfooducate/utils/hash"
	"fastfooducate/utils/token"
	"os"

	"fastfooducate/module/feature/auth/handler"
	"fastfooducate/module/feature/auth/repository"
	"fastfooducate/module/feature/auth/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	repo domain.AuthRepositoryInterface
	serv domain.AuthServiceInterface
	hand domain.AuthHandlerInterface
	hash utils.HashInterface
	jwt  token.JWTInterface
)

func InitializeAuth(db *gorm.DB) {
	secret := os.Getenv("SECRET")
	hash = utils.NewHash()
	jwt = token.NewJWT(secret)

	repo = repository.NewAuthRepository(db)
	serv = service.NewAuthService(repo, hash, jwt)
	hand = handler.NewAuthHandler(serv)
}

func SetupRoutesAuth(app *fiber.App) {
	api := app.Group("/api/v1/auth")
	api.Post("/login", hand.Login)
	api.Post("/register", hand.Register)
}
