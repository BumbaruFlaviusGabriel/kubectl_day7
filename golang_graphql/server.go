package main

import (
	"fmt"
	"golang_graphql/config"
	"golang_graphql/db"
	"golang_graphql/graph"
	"golang_graphql/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	godotenv.Load("app.env")
	config.InitConfig()
	//port := os.Getenv("PORT")
	fmt.Println(os.Getenv("DATABASE_URL"))
	db.InitDatabase()
	router := chi.NewRouter()
	router.Handle("/"+config.GetConfig().APPVersion+"/query", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/"+config.GetConfig().APPVersion+"/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8081")
	log.Fatal(http.ListenAndServe(":"+config.GetConfig().Port, router)) // can use ":"+port too
}
