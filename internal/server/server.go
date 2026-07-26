package server

import (
	"log"
	"net/http"
	"time"

	"github.com/RGBFox/Final-6-sprint/internal/handlers"
)

type Server struct {
	Log      *log.Logger
	Htserver *http.Server
}

// Route - набор настроек для сервера
func Route(l *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.MainHandle)
	mux.HandleFunc("/upload", handlers.UploadHandle)

	myserver := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Log:      l,
		Htserver: &myserver,
	}
}
