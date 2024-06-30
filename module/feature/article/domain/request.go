package domain

type CreateArticleRequest struct {
	Title   string `form:"title" validate:"required"`
	Photo   string `form:"photo"`
	Content string `form:"content" validate:"required"`
}

type UpdateArticleRequest struct {
	Title   string `form:"title"`
	Photo   string `form:"photo"`
	Content string `form:"content"`
}
