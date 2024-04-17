package utils

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2/log"
	"strconv"
)

func ConvStringToUint64(str string) uint64 {
	result, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		log.Errorf("ConvStringToUint64 err: %v", err)
	}

	return result
}

// EncodeBase64 - кодирует в base64
// Входные параметры: данные для кодирования
// Выходные параметры: закодированные данные
func EncodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// DecodeBase64 - декодирует из base64
// Входные параметры: закодированные данные
// Выходные параметры: декодированные данные, ошибка
func DecodeBase64(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
