package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/sicecep/gql-demo/app/database"
	"github.com/sicecep/gql-demo/app/repository"
	"github.com/sicecep/gql-demo/graph"
)

const defaultPort = "8080"

func main() {
	godotenv.Load()
	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := database.NewConnection(config)
	if err != nil {
		panic(err)
	}
	database.Migrate(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	repo := repository.NewBookService(db)
	countryRepo := repository.NewCountryService(db)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		BookRepository:    repo,
		CountryRepository: countryRepo,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
