package services

import (
	"github.com/ky0yk/go-blog/models"
	"github.com/ky0yk/go-blog/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, nil
	}
	
	return newComment, nil
}