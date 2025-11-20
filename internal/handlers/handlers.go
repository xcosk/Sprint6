package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// HandleIndex возвращает HTML файл index.html для корневого эндпоинта
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "index.html")
}

// HandleUpload обрабатывает загрузку файла, конвертирует его содержимое и сохраняет результат
func HandleUpload(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
			return
		}

		// Парсим форму
		err := r.ParseMultipartForm(10 << 20) // 10 MB
		if err != nil {
			logger.Printf("Ошибка парсинга формы: %v", err)
			http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
			return
		}

		// Получаем файл из формы
		file, handler, err := r.FormFile("myFile")
		if err != nil {
			logger.Printf("Ошибка получения файла: %v", err)
			http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Читаем данные из файла
		data, err := io.ReadAll(file)
		if err != nil {
			logger.Printf("Ошибка чтения файла: %v", err)
			http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
			return
		}

		// Конвертируем данные
		converted, err := service.Convert(string(data))
		if err != nil {
			logger.Printf("Ошибка конвертации: %v", err)
			http.Error(w, fmt.Sprintf("Ошибка конвертации: %v", err), http.StatusInternalServerError)
			return
		}

		// Создаем локальный файл с результатом
		// Используем время для генерации имени файла
		fileName := time.Now().UTC().String()
		// Заменяем недопустимые символы в имени файла
		fileName = strings.ReplaceAll(fileName, ":", "-")
		fileName = strings.ReplaceAll(fileName, " ", "_")
		// Получаем расширение из оригинального файла
		ext := filepath.Ext(handler.Filename)
		if ext == "" {
			ext = ".txt"
		}
		fileName = fileName + ext

		// Создаем и записываем файл
		outFile, err := os.Create(fileName)
		if err != nil {
			logger.Printf("Ошибка создания файла: %v", err)
			http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
			return
		}
		defer outFile.Close()

		_, err = outFile.WriteString(converted)
		if err != nil {
			logger.Printf("Ошибка записи в файл: %v", err)
			http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
			return
		}

		// Возвращаем результат конвертации
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, converted)
	}
}
