package middlewares

import (
	"log"
	"net/http"
	"time"
)

type (
	// struct for holding response details
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter // compose original http.ResponseWriter
		responseData        *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b) // write response using original http.ResponseWriter
	r.responseData.size += size            // capture size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode) // write status code using original http.ResponseWriter
	r.responseData.status = statusCode       // capture status code
}

func LoggingMiddleware(next http.Handler) http.HandlerFunc {
	loggingFn := func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lrw := loggingResponseWriter{
			ResponseWriter: rw,
			responseData:   responseData,
		}
		next.ServeHTTP(&lrw, req)

		duration := time.Since(start)

		log.Printf("Request: %s %s %s %s %d %d %s",
			start.Format("2006-01-02 15:04:05"),
			req.Method,
			req.RequestURI,
			req.Proto,
			responseData.status,
			responseData.size,
			duration,
		)

	}
	return http.HandlerFunc(loggingFn)
}
