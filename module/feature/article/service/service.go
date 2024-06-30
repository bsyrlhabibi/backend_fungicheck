package service

import (
	"errors"
	"fastfooducate/module/entities"
	"fastfooducate/module/feature/article/domain"
	"math"
)

type ArticleService struct {
	repo domain.ArticleRepositoryInterface
}

func NewArticleService(repo domain.ArticleRepositoryInterface) domain.ArticleServiceInterface {
	return &ArticleService{
		repo: repo,
	}
}

func (s *ArticleService) CreateArticle(articleData *entities.ArticleModels) (*entities.ArticleModels, error) {
	value := &entities.ArticleModels{
		Title:   articleData.Title,
		Photo:   articleData.Photo,
		Content: articleData.Content,
		Author:  "FASTFOODUCATE",
	}
	createdArticle, err := s.repo.CreateArticle(value)
	if err != nil {
		return nil, errors.New("gagal menambahkan artikel")
	}

	return createdArticle, nil
}

func (s *ArticleService) UpdateArticleById(id uint64, updatedArticle *entities.ArticleModels) (*entities.ArticleModels, error) {
	existingArticle, err := s.repo.GetArticleById(id)
	if err != nil {
		return nil, errors.New("artikel tidak ditemukan")
	}

	if existingArticle == nil {
		return nil, errors.New("artikel tidak ditemukan")
	}

	_, err = s.repo.UpdateArticleById(id, updatedArticle)
	if err != nil {
		return nil, errors.New("gagal mengubah artikel")
	}

	getUpdatedArticle, err := s.repo.GetArticleById(id)
	if err != nil {
		return nil, errors.New("gagal mengambil artikel")
	}

	return getUpdatedArticle, nil
}

func (s *ArticleService) GetArticleById(id uint64) (*entities.ArticleModels, error) {
	result, err := s.repo.GetArticleById(id)
	if err != nil {
		return nil, errors.New("artikel tidak ditemukan")
	}

	return result, nil
}

func (s *ArticleService) GetAllArticles(page, pageSize int) ([]*entities.ArticleModels, int64, error) {
	result, err := s.repo.GetPaginatedArticles(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalItems()
	if err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

func (s *ArticleService) GetArticlePage(currentPage, pageSize, totalItems int) (int, int, int, error) {
	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	nextPage := currentPage + 1
	prevPage := currentPage - 1

	if nextPage > totalPages {
		nextPage = 0
	}

	if prevPage < 1 {
		prevPage = 0
	}

	return totalPages, nextPage, prevPage, nil
}

func (s *ArticleService) DeleteArticleById(id uint64) error {
	schedule, err := s.repo.GetArticleById(id)
	if err != nil {
		return errors.New("schedule not found")
	}

	err = s.repo.DeleteArticle(schedule.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ArticleService) GetAll() ([]*entities.ArticleModels, error) {
	articles, err := s.repo.FindAll()
	if err != nil {
		return nil, errors.New("artikel tidak ditemukan")
	}

	return articles, nil
}

func (s *ArticleService) GetArticlesByTitle(title string) ([]*entities.ArticleModels, error) {
	articles, err := s.repo.FindByTitle(title)
	if err != nil {
		return nil, errors.New("artikel tidak ditemukan")
	}

	return articles, nil
}

func (s *ArticleService) GetAllArticleUser(page, perPage int) ([]*entities.ArticleModels, int64, error) {
	result, err := s.repo.FindAllArticle(page, perPage)
	if err != nil {
		return nil, 0, err
	}
	totalItems, err := s.repo.GetTotalArticleCount()
	if err != nil {
		return nil, 0, err
	}
	return result, totalItems, nil
}

func (s *ArticleService) CalculatePaginationValues(page int, totalItems int, perPage int) (int, int) {
	if page <= 0 {
		page = 1
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(perPage)))
	if page > totalPages {
		page = totalPages
	}

	return page, totalPages
}

func (s *ArticleService) GetNextPage(currentPage int, totalPages int) int {
	if currentPage < totalPages {
		return currentPage + 1
	}

	return totalPages
}

func (s *ArticleService) GetPrevPage(currentPage int) int {
	if currentPage > 1 {
		return currentPage - 1
	}

	return 1
}
