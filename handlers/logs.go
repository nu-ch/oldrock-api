package middlewares

import (
	"net/http"
	"time"

	"github.com/fatih/color"
)

func LogRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		color.Cyan("%s - [Router] %s %s %s %s\n", time.Now().Format(time.RFC3339), r.Host, r.Method, r.URL, r.UserAgent())
		handler.ServeHTTP(w, r)
	})
}

func Logger(msg string) {
	color.Green("%s - [Logger] %s\n", time.Now().Format(time.RFC3339), msg)
}
