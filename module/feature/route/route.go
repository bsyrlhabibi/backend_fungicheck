package route

import (
	"fungicheck/module/feature/article"
	"fungicheck/module/feature/auth"
	users "fungicheck/module/feature/user"
	user "fungicheck/module/feature/user/domain"
	"fungicheck/utils/token"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, jwt token.JWTInterface, userService user.UserServiceInterface) {
	users.InitializeUser(db)
	users.SetupRoutesUser(app)
	auth.InitializeAuth(db)
	auth.SetupRoutesAuth(app)
	article.InitializeArticle(db)
	article.SetupRoutesArticle(app, jwt, userService)

	// Jika Anda ingin mengaktifkan schedule, tambahkan kode berikut
	// schedule.InitializeSchedule(db)
	// schedule.SetupRoutesSchedule(app, jwt, userService)
}
