package services

import "github.com/ky0yk/go-blog/models"

type ArticleService interface {
 PostArticleService (article models.Article) (models.Article, error)
 GetArticleListService  (page int) ([]models.Article, error)
 GetArticleService (articleID int) (models.Article, error)
 PostNiceService (article models.Article) (models.Article, error)
}


type CommentService interface {
	PostCommentService (comment models.Comment) (models.Comment, error)
}


