package handler

import (
	"fungicheck/module/entities"
	"fungicheck/module/feature/article/domain"
	"fungicheck/utils/response"
	"fungicheck/utils/upload"
	"fungicheck/utils/validator"
	"mime/multipart"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ArticleHandler struct {
	service domain.ArticleServiceInterface
}

func NewArticleHandler(service domain.ArticleServiceInterface) domain.ArticleHandlerInterface {
	return &ArticleHandler{
		service: service,
	}
}

func (h *ArticleHandler) CreateArticle(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}

	articleRequest := new(domain.CreateArticleRequest)
	file, err := c.FormFile("photo")
	var uploadedURL string
	if err == nil {
		fileToUpload, err := file.Open()
		if err != nil {
			return response.SendStatusInternalServerResponse(c, "Gagal membuka file: "+err.Error())
		}

		defer func(fileToUpload multipart.File) {
			_ = fileToUpload.Close()
		}(fileToUpload)

		uploadedURL, err = upload.ImageUploadHelper(fileToUpload)
		if err != nil {
			return response.SendStatusInternalServerResponse(c, "Gagal mengunggah foto: "+err.Error())
		}
	}

	if err := c.BodyParser(articleRequest); err != nil {
		return response.SendBadRequestResponse(c, "Format input yang Anda masukkan tidak sesuai: "+err.Error())
	}

	if err := validator.ValidateStruct(articleRequest); err != nil {
		return response.SendBadRequestResponse(c, "Validasi gagal: "+err.Error())
	}

	newArticle := &entities.ArticleModels{
		Title:   articleRequest.Title,
		Photo:   uploadedURL,
		Content: articleRequest.Content,
	}

	createdArticle, err := h.service.CreateArticle(newArticle)
	if err != nil {
		return response.SendStatusInternalServerResponse(c, "Gagal menambahkan artikel: "+err.Error())
	}

	return response.SendStatusCreatedResponse(c, "Berhasil menambahkan artikel", domain.FormatArticle(createdArticle))
}

func (h *ArticleHandler) GetAllArticles(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page"))
	pageConv, _ := strconv.Atoi(strconv.Itoa(page))
	perPage := 8

	var articles []*entities.ArticleModels
	var totalItems int64
	var err error

	search := c.Query("search")

	if search != "" {
		articles, err = h.service.GetArticlesByTitle(search)
	} else {
		articles, totalItems, err = h.service.GetAllArticleUser(pageConv, perPage)
	}

	if err != nil {
		return response.SendStatusInternalServerResponse(c, "Failed to get article list: "+err.Error())
	}

	var activeArticles []*entities.ArticleModels
	for _, article := range articles {
		if article.DeletedAt == nil {
			activeArticles = append(activeArticles, article)
		}
	}

	currentPage, totalPages := h.service.CalculatePaginationValues(pageConv, int(totalItems), perPage)
	nextPage := h.service.GetNextPage(currentPage, totalPages)
	prevPage := h.service.GetPrevPage(currentPage)

	return response.PaginationBuildResponse(
		c,
		fiber.StatusOK,
		"Success get article list",
		domain.FormatterArticle(activeArticles),
		currentPage,
		int(totalItems),
		totalPages,
		nextPage,
		prevPage,
	)
}
