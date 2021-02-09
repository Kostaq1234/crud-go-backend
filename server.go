package main

import (
	customMiddleware "github.com/Kostaq1234/graphql/middleware"
	"github.com/Kostaq1234/graphql/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Kostaq1234/graphql/graph"
	"github.com/Kostaq1234/graphql/graph/generated"
)

const defaultPort = "8080"

func main() {
	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "Kostaq",
		Database: "graphqlTest",
	})
	defer DB.Close()
	//DB.AddQueryHook(postgres.DBLogger{})
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userRepo := postgres.UserRepo{DB: DB}
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		Debug:            true,
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:8080"},
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(userRepo))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{MeetupsRepo: postgres.MeetupsRepo{DB: DB}, UserRepo: userRepo}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graph.DataloaderMiddleware(DB, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
