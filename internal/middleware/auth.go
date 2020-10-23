package middleware

import (
	db "HowToGraphql/db/sqlc"
	"HowToGraphql/graph/model"
	"HowToGraphql/pkg/jwt"
	"context"
	"log"
	"net/http"
	"strconv"
)

type userContextKey string

var (
	key = userContextKey("user")
	testKey = userContextKey("testKey")
)

func Middleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// log.Println("Authenticating")

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

			user, err := db.DbQueries.GetUser(r.Context(), username)
			if err != nil {
				log.Println("Invalid Token", err)
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			qlUser := &model.User{
				ID:   strconv.FormatInt(user.ID, 10),
				Name: user.Username,
			}

			// log.Println("Adding user to context")
			ctx := context.WithValue(r.Context(), key, qlUser)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	log.Println(ctx)
	raw := ctx.Value(key)
	if raw == nil {
		log.Println("Cannot find key")
		return nil
	}
	return raw.(*model.User)
}