package services

import (
	"github.com/ky0yk/go-blog/models"
	"github.com/ky0yk/go-blog/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, nil
	}
	
	return newComment, nil
}