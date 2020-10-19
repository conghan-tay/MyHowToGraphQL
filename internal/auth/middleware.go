package auth

import (
	db "HowToGraphql/db/sqlc"
	"HowToGraphql/graph/model"
	"HowToGraphql/pkg/jwt"
	"context"
	"log"
	"net/http"
	"strconv"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Authenticating")

		header := r.Header.Get("Authorization")

		if header == "" {
			next.ServeHTTP(w, r)
			return
		}

		tokenStr := header
		username, err := jwt.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		user, err :=db.DbQueries.GetUser(r.Context(), username)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		qlUser := model.User{
			ID:   strconv.FormatInt(user.ID, 10),
			Name: user.Username,
		}

		ctx := context.WithValue(r.Context(), userCtxKey, qlUser)

		r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}