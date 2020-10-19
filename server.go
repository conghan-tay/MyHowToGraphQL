package main

import (
	db "HowToGraphql/db/sqlc"
	"HowToGraphql/graph"
	"HowToGraphql/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "9080"

func middlewareplaygroundHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing playgroundHandler")
		next.ServeHTTP(w, r)
	})
}

func middlewarequery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing query")
		next.ServeHTTP(w, r)
	})
}


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db.InitDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", middlewareplaygroundHandler(playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", middlewarequery(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
