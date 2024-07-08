package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bananichdev/ozon-graphql-api/db"
	"github.com/bananichdev/ozon-graphql-api/graph"
	"github.com/bananichdev/ozon-graphql-api/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if len(os.Args) >= 2 {
		mode := os.Args[1]
		if mode == "-inmemory" {
			settings.MemoryMode = true
		} else if mode == "-db" {
			settings.MemoryMode = false
		} else {
			panic("wrong args\nargs:\n-inmemory: everything is stored in memory\n-db: everything is stored in database\ndb mode is by default")
		}
	}

	port := settings.Port

	var srv *handler.Server

	if settings.MemoryMode {
		srv = handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	} else {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", settings.DBHost, settings.DBUser, settings.DBPass, settings.DBName, settings.DBPort)
		DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		srv = handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
			PostRepo:    db.PostRepo{DB: DB},
			CommentRepo: db.CommentRepo{DB: DB},
		}}))
	}

	srv.AddTransport(&transport.Websocket{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
