package middlewares

import (
	"net/http"
	"runtime"
	"time"

	"github.com/fatih/color"
)

func LogRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		color.Cyan("%s - [Router] %s %s %s %s\n", time.Now().Format(time.RFC3339), r.Host, r.Method, r.URL, r.UserAgent())
		handler.ServeHTTP(w, r)
	})
}

func DebugLogger(msg string) {
	color.Green("%s - [Logger] %s\n", time.Now().Format(time.RFC3339), msg)
}

func ErrorLogger(err error) {
	pc, _, line, _ := runtime.Caller(1)
	color.Red("%s - [Exception] %s::%s:line:%d\n", time.Now().Format(time.RFC3339), err, runtime.FuncForPC(pc).Name(), line)

}
