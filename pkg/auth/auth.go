package auth

import (
	"errors"
	"fmt"
	"refresh/internal/app/config"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Token struct{}

func (t *Token) ParseToken(tokenString, secret string) (string, error) {
	if len(tokenString) >= 7 && tokenString[0:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	return id, nil
}

func (t *Token) createToken(id string, expire int, tokenSecret string) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Minute * time.Duration(expire)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	acc, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}

	return acc, nil
}

func (t *Token) CreateAccessToken(id string) (string, error) {
	return t.createToken(id, config.InitConfig().ACCES_EXPIRE, config.InitConfig().ACCES_TOKEN)
}

func (t *Token) CreateRefreshToken(id string) (string, error) {
	return t.createToken(id, config.InitConfig().REFRESH_EXPIRE, config.InitConfig().REFRESH_TOKEN)
}

func (t *Token) ExtractToken(c *fiber.Ctx) (string, error) {

	user := c.Locals("user")
	if user == nil {
		return "", errors.New("invalid token")
	}

	claims := user.(fiber.Map)

	id, ok := claims["id"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}

	return id, nil
}
