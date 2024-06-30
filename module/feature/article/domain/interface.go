package domain

import (
	"fastfooducate/module/entities"

	"github.com/gofiber/fiber/v2"
)

type ArticleRepositoryInterface interface {
	CreateArticle(req *entities.ArticleModels) (*entities.ArticleModels, error)
	UpdateArticleById(id uint64, updatedArticle *entities.ArticleModels) (*entities.ArticleModels, error)
	GetArticleById(id uint64) (*entities.ArticleModels, error)
	GetTotalItems() (int64, error)
	GetPaginatedArticles(page, pageSize int) ([]*entities.ArticleModels, error)
	DeleteArticle(article uint64) error
	FindAll() ([]*entities.ArticleModels, error)
	FindByTitle(title string) ([]*entities.ArticleModels, error)
	FindAllArticle(page, perPage int) ([]*entities.ArticleModels, error)
	GetTotalArticleCount() (int64, error)
}

type ArticleServiceInterface interface {
	CreateArticle(req *entities.ArticleModels) (*entities.ArticleModels, error)
	UpdateArticleById(id uint64, updatedArticle *entities.ArticleModels) (*entities.ArticleModels, error)
	GetAllArticles(page, pageSize int) ([]*entities.ArticleModels, int64, error)
	GetArticlePage(currentPage, pageSize, totalItems int) (int, int, int, error)
	GetArticleById(id uint64) (*entities.ArticleModels, error)
	DeleteArticleById(id uint64) error
	GetNextPage(currentPage int, totalPages int) int
	GetPrevPage(currentPage int) int
	CalculatePaginationValues(page int, totalItems int, perPage int) (int, int)
	GetAllArticleUser(page, perPage int) ([]*entities.ArticleModels, int64, error)
	GetAll() ([]*entities.ArticleModels, error)
	GetArticlesByTitle(title string) ([]*entities.ArticleModels, error)
}

type ArticleHandlerInterface interface {
	CreateArticle(c *fiber.Ctx) error
	// UpdateArticleById(c *fiber.Ctx) error
	GetAllArticles(c *fiber.Ctx) error
	// GetArticleById(c *fiber.Ctx) error
	// DeleteArticleById(c *fiber.Ctx) error
}
