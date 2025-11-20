package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server представляет HTTP-сервер
type Server struct {
	logger *log.Logger
	server *http.Server
}

// New создает новый HTTP-сервер с роутером и регистрирует хендлеры
func New(logger *log.Logger) *Server {
	// Создаем HTTP-роутер
	mux := http.NewServeMux()

	// Регистрируем хендлеры
	mux.HandleFunc("/", handlers.HandleIndex)
	mux.HandleFunc("/upload", handlers.HandleUpload(logger))

	// Создаем экземпляр http.Server
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		server: httpServer,
	}
}

// Start запускает HTTP-сервер
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
