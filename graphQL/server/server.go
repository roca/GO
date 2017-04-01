package server

import (
	"math/rand"
	"net/http"
	"os"

	"github.com/GOCODE/graphQL/handlers"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var Version int

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello World!", nil
			},
		},
		"randomNumber": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return rand.Intn(100), nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func Start() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/graphiql", handlers.IndexPage).Methods("GET")

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	// serve a GraphQL endpoint at `/graphql`

	http.Handle("/graphql", h)
	http.Handle("/graphiql", gorillaRoute)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	port := os.Getenv("PORT")

	// and serve!
	http.ListenAndServe(":"+port, nil)
}

// curl -XPOST http://local.rit.aws.regeneron.com:8080/graphql -H 'Content-Type: application/graphql' -d 'query Root{ latestPost }'
