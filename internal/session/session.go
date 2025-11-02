package session

import (
	"RIP/internal/models"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "your-secret-key" // В продакшене хранить в переменной окружения

type UserSession struct {
	UserID      uint
	IsModerator bool
}

type Claims struct {
	UserID      uint `json:"user_id"`
	IsModerator bool `json:"is_moderator"`
	jwt.StandardClaims
}

func CreateSession(w http.ResponseWriter, user *models.User) string {
	claims := &Claims{
		UserID:      user.ID,
		IsModerator: user.IsModerator,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return ""
	}
	w.Header().Set("Authorization", "Bearer "+tokenString)
	return tokenString
}

func GetUser(r *http.Request) *UserSession {
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return nil
	}
	return &UserSession{UserID: claims.UserID, IsModerator: claims.IsModerator}
}

func DestroySession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Authorization", "")
}
