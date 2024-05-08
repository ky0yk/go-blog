package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ky0yk/go-blog/apperrors"
	"github.com/ky0yk/go-blog/controllers/services"
	"github.com/ky0yk/go-blog/models"
)

type CommentController struct {
	services services.CommentService
}

func NewCommentController(s services.CommentService) * CommentController{
	return &CommentController{services: s}
}


// POST /comment のハンドラ
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodefailed.Wrap(err, "bad request body")
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment, err := c.services.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}