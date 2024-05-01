package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ky0yk/go-blog/models"
)

func HelloHandler (w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler (w http.ResponseWriter, req *http.Request){

	// 1. バイトスライス reqBodybuffer を何らかの形で用意
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "cannot get content lenght\n", http.StatusBadRequest)
		return
	}
	reqBodyBuffer := make([]byte, length)

	// 2. Read メソッドでリクエストボディを読み出し
	if _, err := req.Body.Read(reqBodyBuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
	}

	// 3. ボディを Close する
	defer req.Body.Close()

	var reqArticle models.Article
	if err := json.Unmarshal(reqBodyBuffer, &reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article := reqArticle
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w. Write(jsonData)
}

func ArticListleHandler (w http.ResponseWriter, req *http.Request){
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (page %d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w. Write(jsonData)
}

func ArticleDetailHandler (w http.ResponseWriter, req *http.Request){
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (articleId %d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w. Write(jsonData)
}

func PostNiceHandler (w http.ResponseWriter, req *http.Request){
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w. Write(jsonData)
}

func PostCommentHandler (w http.ResponseWriter, req *http.Request){
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w. Write(jsonData)
}