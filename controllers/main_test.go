package controllers_test

import (
	"testing"

	"github.com/ky0yk/go-blog/controllers"
	"github.com/ky0yk/go-blog/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M){

	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}