package repository

import (
	"fastfooducate/module/entities"
	"fastfooducate/module/feature/article/domain"
	"time"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) domain.ArticleRepositoryInterface {
	return &ArticleRepository{
		db: db,
	}
}

func (r *ArticleRepository) CreateArticle(req *entities.ArticleModels) (*entities.ArticleModels, error) {
	if err := r.db.Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (r *ArticleRepository) UpdateArticleById(id uint64, updatedArticle *entities.ArticleModels) (*entities.ArticleModels, error) {
	var article *entities.ArticleModels
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&article).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(article).Updates(updatedArticle).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (r *ArticleRepository) DeleteArticle(id uint64) error {
	article := &entities.ArticleModels{}
	if err := r.db.First(article, id).Error; err != nil {
		return err
	}

	if err := r.db.Model(article).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func (r *ArticleRepository) GetTotalItems() (int64, error) {
	var totalItems int64

	if err := r.db.Where("deleted_at IS NULL").
		Model(&entities.ArticleModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *ArticleRepository) GetPaginatedArticles(page, pageSize int) ([]*entities.ArticleModels, error) {
	var articles []*entities.ArticleModels

	offset := (page - 1) * pageSize

	if err := r.db.Where("deleted_at IS NULL").
		Offset(offset).Limit(pageSize).Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *ArticleRepository) GetArticleById(id uint64) (*entities.ArticleModels, error) {
	var article *entities.ArticleModels

	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (r *ArticleRepository) FindAll() ([]*entities.ArticleModels, error) {
	var articles []*entities.ArticleModels
	err := r.db.Where("deleted_at IS NULL").Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *ArticleRepository) FindByTitle(title string) ([]*entities.ArticleModels, error) {
	var articles []*entities.ArticleModels
	err := r.db.Where("deleted_at IS NULL AND title LIKE?", "%"+title+"%").Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *ArticleRepository) GetTotalArticleCount() (int64, error) {
	var count int64
	err := r.db.Model(&entities.ArticleModels{}).Where("deleted_at IS NULL").Count(&count).Error
	return count, err
}

func (r *ArticleRepository) FindAllArticle(page, perPage int) ([]*entities.ArticleModels, error) {
	var articles []*entities.ArticleModels
	offset := (page - 1) * perPage

	err := r.db.Where("deleted_at IS NULL").Offset(offset).Limit(perPage).Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}
