package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Создание JWT-токена
var jwtKey = []byte("my_secret_key")

// Структура для сохранения логина и пароля
type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// Структура для хранения информации о пользователе в токене
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// Функция для генерации JWT
func GenerateJWT(username string) (string, error) {
    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// Создание маршрута для аутентификации (логин)
func Login(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    // Предположим, что данные логина передаются в JSON
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Проверка данных (в реальности ты бы проверял с базой данных)
    if creds.Username != "user" || creds.Password != "password" {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // Генерация JWT
    token, err := GenerateJWT(creds.Username)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // Возвращаем токен клиенту
    http.SetCookie(w, &http.Cookie{
        Name:    "token",
        Value:   token,
        Expires: time.Now().Add(5 * time.Minute),
    })
}


// Маршрут с проверкой токена (защищенный ресурс
func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
    // Извлекаем токен из куки
    cookie, err := r.Cookie("token")
    if err != nil {
        if err == http.ErrNoCookie {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    tokenStr := cookie.Value

    claims := &Claims{}

    // Проверяем токен
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    if !token.Valid {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // Если токен валиден, можно возвращать защищенные данные
    fmt.Fprintf(w, "Welcome %s!", claims.Username)
}
