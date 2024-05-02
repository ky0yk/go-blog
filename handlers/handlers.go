package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ky0yk/go-blog/models"
)

func HelloHandler (w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler (w http.ResponseWriter, req *http.Request){

	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article := reqArticle

	json.NewEncoder(w).Encode(article)

}

func ArticListleHandler (w http.ResponseWriter, req *http.Request){
	queryMap := req.URL.Query()

	// クエリパラメータ page を取得
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

	// 暫定対応: 「変数 page が使われていない」というコンパイルエラーを回避
	log.Println(page)


	articleList := []models.Article{models.Article1, models.Article2}

	if err := json.NewEncoder(w).Encode(articleList); err != nil {
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}
}

func ArticleDetailHandler (w http.ResponseWriter, req *http.Request){
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	// 暫定対応: 「変数 articleID が使われていない」というコンパイルエラーを回避
	log.Println(articleID)

	article := models.Article1

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
}

func PostNiceHandler (w http.ResponseWriter, req *http.Request){
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article := reqArticle

	json.NewEncoder(w).Encode(article)

}

func PostCommentHandler (w http.ResponseWriter, req *http.Request){
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	comment := reqComment

	json.NewEncoder(w).Encode(comment)
}