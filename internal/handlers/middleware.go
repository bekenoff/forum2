package handlers

import (
	"log"
	"net/http"
	"os"
)

var l = NewLogger()

func (h *Handler) AllHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.Info(r.Method + " successfully working")
		next.ServeHTTP(w, r)
	})
}

type logger struct {
	info *log.Logger
	e    *log.Logger
}

func NewLogger() Logger {
	return &logger{
		info: log.New(os.Stdout, "", 0),
		e:    log.New(os.Stderr, "", 0),
	}
}

type Logger interface {
	Info(msg string)
	Error(msg string, err error)
}

func (l *logger) Info(msg string) {
	l.info.Println("INFO:", "\033[0;34m")
}

func (l *logger) Error(msg string, err error) {
	l.e.Println("ERROR:", "\033[0;31m")
}
