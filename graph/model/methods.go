package model

import (
	db "HowToGraphql/db/sqlc"
	"HowToGraphql/internal/auth"
	"golang.org/x/net/context"
	"log"
)

func (login *Login) Authenticate(ctx context.Context) bool {
	result, err := db.DbQueries.GetUser(ctx, login.Username)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return auth.CheckPasswordHash(login.Password, result.Password)
}