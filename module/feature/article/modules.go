package article

import (
	"fungicheck/module/feature/article/handler"
	"fungicheck/module/feature/article/repository"
	"fungicheck/module/feature/article/service"
	"fungicheck/module/feature/middleware"
	"fungicheck/utils/token"

	"fungicheck/module/feature/article/domain"
	user "fungicheck/module/feature/user/domain"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	repo domain.ArticleRepositoryInterface
	serv domain.ArticleServiceInterface
	hand domain.ArticleHandlerInterface
)

func InitializeArticle(db *gorm.DB) {
	repo = repository.NewArticleRepository(db)
	serv = service.NewArticleService(repo)
	hand = handler.NewArticleHandler(serv)
}

func SetupRoutesArticle(app *fiber.App, jwt token.JWTInterface, userService user.UserServiceInterface) {
	api := app.Group("/api/v1/article")
	api.Get("/all", hand.GetAllArticles)
	api.Post("/create", middleware.AuthMiddleware(jwt, userService), hand.CreateArticle)
}
