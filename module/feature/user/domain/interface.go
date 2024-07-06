package domain

import (
	"fungicheck/module/entities"

	"github.com/gofiber/fiber/v2"
)

type UserRepositoryInterface interface {
	GetUserByID(userID uint64) (*entities.UserModels, error)
}

type UserServiceInterface interface {
	GetUserByID(userID uint64) (*entities.UserModels, error)
}

type UserHandlerInterface interface {
	GetUserByID(c *fiber.Ctx) error
}
