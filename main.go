package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	mySigningKey := []byte("MySecretKey")

	// Создание claims (данных, которые будут содержаться в токене)
	claims := jwt.MapClaims{
		"username": "somnium11",
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Токен истекает через 1 час
	}

	// Создание токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписание токена секретным ключом
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("Error creating token:", err)
		return
	}

	fmt.Println("Generated JWT token:", tokenString)
}
