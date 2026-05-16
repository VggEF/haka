package utils

import (
	"time"
)

// GetWeekStart возвращает дату начала недели (понедельник)
func GetWeekStart(date time.Time) time.Time {
	weekday := date.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	}
	offset := int(weekday) - 1
	return date.AddDate(0, 0, -offset)
}

// GetWeekEnd возвращает дату конца недели (воскресенье)
func GetWeekEnd(date time.Time) time.Time {
	return GetWeekStart(date).AddDate(0, 0, 6)
}

// FormatDate форматирует дату в формате DD.MM.YYYY
func FormatDate(date time.Time) string {
	return date.Format("02.01.2006")
}

// FormatDateTime форматирует дату и время
func FormatDateTime(date time.Time) string {
	return date.Format("02.01.2006 15:04")
}

// DaysUntil возвращает количество дней до указанной даты
func DaysUntil(date time.Time) int {
	now := time.Now()
	diff := date.Sub(now)
	return int(diff.Hours() / 24)
}

// IsToday проверяет, является ли дата сегодняшней
func IsToday(date time.Time) bool {
	now := time.Now()
	return date.Year() == now.Year() && date.Month() == now.Month() && date.Day() == now.Day()
}
