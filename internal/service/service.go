package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("пустая строка")
	}

	if isMorse(trimmed) {
		return morse.ToText(trimmed), nil
	}

	return morse.ToMorse(trimmed), nil
}

func isMorse(s string) bool {
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' {
			return false
		}
	}
	return true
}
