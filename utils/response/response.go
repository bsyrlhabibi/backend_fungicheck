package response

import (
	"github.com/gofiber/fiber/v2"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type PaginationResponse struct {
	CurrentPage int `json:"current_page"`
	TotalItems  int `json:"total_items"`
	TotalPages  int `json:"total_pages"`
	PrevPage    int `json:"prev_page"`
	NextPage    int `json:"next_page"`
}

type PaginationData struct {
	Message    string             `json:"message"`
	Data       interface{}        `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

func SuccessBuildResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := SuccessResponse{
		Message: message,
		Data:    data,
	}

	return c.Status(statusCode).JSON(response)
}

func ErrorBuildResponse(c *fiber.Ctx, statusCode int, message string) error {
	response := ErrorResponse{
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

func SuccessBuildWithoutResponse(c *fiber.Ctx, statusCode int, message string) error {
	response := ErrorResponse{
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

func PaginationBuildResponse(c *fiber.Ctx, statusCode int, message string, data interface{}, currentPage, totalItems, totalPages, nextPage, prevPage int) error {
	pagination := PaginationResponse{
		CurrentPage: currentPage,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		PrevPage:    prevPage,
		NextPage:    nextPage,
	}

	paginationData := PaginationData{
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}

	return c.Status(statusCode).JSON(paginationData)
}

func SendStatusForbiddenResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": message})
}

func SendStatusCreatedResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": message, "data": data})
}

func SendStatusOkResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message})
}

func SendStatusOkWithDataResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message, "data": data})
}

func SendStatusOkWithDataResponses(c *fiber.Ctx, message string, extraMessage string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message, "extra_message": extraMessage, "data": data})
}

func SendStatusInternalServerResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": message})
}

func SendBadRequestResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": message})
}

func SendSuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message, "data": data})
}

func SendStatusConflictResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": message})
}

func SendStatusNotFoundResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": message})
}

func SendStatusUnauthorizedResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": message})
}
