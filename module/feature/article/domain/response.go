package domain

import (
	"fungicheck/module/entities"
	"time"
)

type ArticleFormatter struct {
	ID      uint64    `json:"id"`
	Title   string    `json:"title"`
	Photo   string    `json:"photo"`
	Content string    `json:"content"`
	Author  string    `json:"author"`
	Date    time.Time `json:"date"`
}

func FormatArticle(article *entities.ArticleModels) *ArticleFormatter {
	articleFormatter := &ArticleFormatter{}
	articleFormatter.ID = article.ID
	articleFormatter.Title = article.Title
	articleFormatter.Photo = article.Photo
	articleFormatter.Content = article.Content
	articleFormatter.Author = article.Author
	articleFormatter.Date = article.CreatedAt

	return articleFormatter
}

func ResponseArrayArticles(data []*entities.ArticleModels) []*ArticleFormatter {
	res := make([]*ArticleFormatter, 0)

	for _, article := range data {
		articleRes := &ArticleFormatter{
			ID:      article.ID,
			Title:   article.Title,
			Content: article.Content,
			Photo:   article.Photo,
			Author:  article.Author,
		}
		res = append(res, articleRes)
	}

	return res
}

func FormatterArticle(articles []*entities.ArticleModels) []*ArticleFormatter {
	var articleFormatter []*ArticleFormatter

	for _, article := range articles {
		formatArticle := FormatArticle(article)
		articleFormatter = append(articleFormatter, formatArticle)
	}

	return articleFormatter
}
