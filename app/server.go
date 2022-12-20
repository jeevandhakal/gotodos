package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/jeevandhakal/todos/app/db"
	"github.com/jeevandhakal/todos/auth"
	"github.com/jeevandhakal/todos/graph/resolver"
	"github.com/jeevandhakal/todos/pkg/adapter/controller"
	"github.com/jeevandhakal/todos/pkg/registry"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db := db.InitDB()

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	ctrl := newController(db)
	srv := handler.NewDefaultServer(resolver.NewSchema(ctrl))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func newController(db *gorm.DB) controller.Controller {
	r := registry.New(db)
	return r.NewController()
}
