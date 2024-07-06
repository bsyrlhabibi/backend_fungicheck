package domain

import (
	"fungicheck/module/entities"

	"github.com/gofiber/fiber/v2"
)

type AuthRepositoryInterface interface {
	GetUsersByEmail(email string) (*entities.UserModels, error)
	CreateUser(req *entities.UserModels) (*entities.UserModels, error)
}

type AuthServiceInterface interface {
	Login(email, password string) (*entities.UserModels, string, error)
	Register(req *RegisterRequest) (*entities.UserModels, error)
}

type AuthHandlerInterface interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}
