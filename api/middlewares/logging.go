package middlewares

import (
	"log"
	"net/http"
)

type resLogginWriter struct {
	http.ResponseWriter
	code int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLogginWriter {
	return &resLogginWriter{ResponseWriter: w, code: http.StatusOK}
}

func (rsw *resLogginWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleWare (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
		traceID := newTraceID()

		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}