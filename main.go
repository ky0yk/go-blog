package main

import (
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// var (
// 	dbUser     = os.Getenv("DB_USER")
// 	dbPassword = os.Getenv("DB_PASSWORD")
// 	dbDatabase = os.Getenv("DB_NAME")
// 	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
// )

// func main() {
// 	db, err := sql.Open("mysql", dbConn)
// 	if err != nil {
// 		log.Println("fail to connect DB")
// 		return
// 	}

// 	r := api.NewRouter(db)

// 	log.Println("server start at port 8080")
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

func main() {
	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
		io.WriteString(w, "Hello, wordl!\n")
	})

	http.Handle("/", myMiddleware2(myMiddleware1(helloHandler)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myMiddleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "Pre-process1\n")
		next.ServeHTTP(w, r)
		io.WriteString(w, "Post-process1\n")
	})
}

func myMiddleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "Pre-process2\n")
		next.ServeHTTP(w, r)
		io.WriteString(w, "Post-process2\n")
	})
}