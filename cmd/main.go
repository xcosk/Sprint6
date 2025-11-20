package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// Создаем сервер
	srv := server.New(logger)

	// Запускаем сервер
	if err := srv.Start(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
