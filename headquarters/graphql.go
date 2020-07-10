package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type resolver struct{}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func setupServer() {
	s, err := getSchema("./schema.graphql")
	if err != nil {
		log.Fatal("unable to get schema: ", err)
	}

	schema := graphql.MustParseSchema(s, &resolver{}, graphql.UseStringDescriptions())

	http.Handle("/graphql", &relay.Handler{Schema: schema})
}

// StartServer sets up and starts the GraphQL service
func StartServer() {
	setupServer()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
