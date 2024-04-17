package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"idon.com/cfg"
	"strconv"
	"time"
)

var (
	ErrInvalidClaimsType = errors.New("неверный тип данных в JWT")
	ErrInvalidToken      = errors.New("токен не действителен")
)

func MakeAccessJWT(userID uint64) (string, error) {
	claims := &jwt.RegisteredClaims{
		Issuer:    "idon.com/crm",
		Subject:   "user_id",
		Audience:  []string{"client_hola"},
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        strconv.FormatUint(userID, 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.GetAppConfig().SecretKey))
}

// ParseAndValidateJWT - парсит JWT
// Входные параметры: JWT, ключ для подписи JWT
// Выходные параметры: данные из JWT, ошибка
func ParseAndValidateJWT(tokenString string, keyFunc jwt.Keyfunc) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		keyFunc,
		jwt.WithIssuer("idon.com/crm"),
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, ErrInvalidClaimsType
	}
	if !token.Valid {
		return claims, ErrInvalidToken
	}

	return claims, nil
}
