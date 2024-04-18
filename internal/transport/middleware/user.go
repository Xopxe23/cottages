package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
type ctxValue string

const (
	ctxUserID ctxValue = "userId"
)

func AddUserIdInContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := getTokenFromRequest(r)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		userId, err := parseToken(token)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), ctxUserID, userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func getTokenFromRequest(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return headerParts[1], nil
}

// func getTokenFromCookie(r *http.Request) (string, error) {
// 	token, err := r.Cookie("accessToken")
// 	if err != nil {
// 		return "", err
// 	}
// 	return token.Value, err
// }

func parseToken(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("sample secret"), nil
	})
	if err != nil {
		return "", err
	}
	if !t.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	userId, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid subject")
	}

	expAt, ok := claims["expiredAt"].(float64)
	if !ok {
		return "", errors.New("invalid expired")
	}
	if int64(expAt) < time.Now().Unix() {
		return "", errors.New("token expired")
	}
	return userId, nil
}
