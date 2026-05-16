package utils

import (
	"regexp"
	"strings"
)

// ValidateEmail проверяет корректность email
func ValidateEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

// ValidatePhone проверяет корректность телефона
func ValidatePhone(phone string) bool {
	regex := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return regex.MatchString(phone)
}

// ValidateStudentLogin проверяет формат логина студента (23-ПМбо-014)
func ValidateStudentLogin(login string) bool {
	regex := regexp.MustCompile(`^\d{2}-[А-Я]{2}[а-я]{2}-\d{3}$`)
	return regex.MatchString(login)
}

// ValidateGroupName проверяет название группы (23-ПМбо-1)
func ValidateGroupName(group string) bool {
	regex := regexp.MustCompile(`^\d{2}-[А-Я]{2}[а-я]{2}-\d$`)
	return regex.MatchString(group)
}

// TruncateString обрезает строку до указанной длины
func TruncateString(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "..."
}

// Slugify создает URL-дружественную строку
func Slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	s = regexp.MustCompile(`[^a-zа-я0-9-]`).ReplaceAllString(s, "")
	s = regexp.MustCompile(`-+`).ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}
