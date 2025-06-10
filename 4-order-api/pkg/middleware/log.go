package middleware

import (
	"net/http"
	"time"
	log "github.com/sirupsen/logrus"
)


func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode: http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		log.WithFields(log.Fields{
			"status_code": wrapper.StatusCode,
			"method": r.Method,
			"url_path": r.URL.Path,
			"handle_duration": time.Since(start),
		}).Info("Request")
	})
}