package utils

import (
	cryptorand "crypto/rand"
	"encoding/hex"
	"math/rand"
)

// RandomString генерирует случайную строку заданной длины
func RandomString(length int) string {
	bytes := make([]byte, length/2)
	cryptorand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// RandomInt генерирует случайное число в диапазоне
func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// RandomBool генерирует случайное булево значение
func RandomBool() bool {
	return rand.Intn(2) == 1
}
