package route

import (
	"fastfooducate/module/feature/article" // Pastikan ini sesuai dengan struktur folder Anda
	"fastfooducate/module/feature/auth"    // Pastikan ini sesuai dengan struktur folder Anda
	users "fastfooducate/module/feature/user"
	user "fastfooducate/module/feature/user/domain"
	"fastfooducate/utils/token"

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
