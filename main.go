package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	handler "github.com/graphql-go/handler"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryType,
		Mutation: MutationType,
	})
	if err != nil {
		log.Fatal(err)
	}
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/go_graphql")
	if err != nil {
		log.Fatal(err)
	}

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/graphql", h)
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
