package main

import (
	db "HowToGraphql/db/sqlc"
	"HowToGraphql/graph"
	"HowToGraphql/graph/generated"
	"HowToGraphql/internal/middleware"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "9080"

//func middlewarequery(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Println("Executing query")
//		next.ServeHTTP(w, r)
//	})
//}


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//router := chi.NewRouter()
	//router.Use(middleware.Middleware())
	db.InitDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.Middleware(srv))
	//router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
