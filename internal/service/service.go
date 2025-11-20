package service

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert автоматически определяет и конвертирует строку
func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", fmt.Errorf("пустая строка")
	}

	// Определяем тип строки
	isMorse, err := autoDetect(input)
	if err != nil {
		return "", err
	}

	if isMorse {
		// Конвертируем код Морзе в текст
		return morse.ToText(input), nil
	} else {
		// Конвертируем текст в код Морзе
		return morse.ToMorse(input), nil
	}
}

// autoDetect определяет, является ли строка кодом Морзе или обычным текстом
// Возвращает true, если это код Морзе, false - если текст
func autoDetect(input string) (bool, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return false, fmt.Errorf("пустая строка")
	}

	var dots, dashes, letters int
	for _, r := range input {
		switch {
		case r == '.':
			dots++
		case r == '-':
			dashes++
		case unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsPunct(r):
			letters++
		case r == ' ' || r == '\n' || r == '\t':
			// Пробельные символы пропускаем
		}
	}

	hasMorse := dots > 0 || dashes > 0
	hasText := letters > 0

	switch {
	case hasMorse && !hasText:
		return true, nil // Это код Морзе
	case hasText:
		return false, nil // Это текст
	default:
		return false, fmt.Errorf("не удалось определить тип")
	}
}
