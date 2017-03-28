package server

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var Version int

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p types.ResolveParams) interface{} {
				return "Hello World!"
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func Start() {

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	// and serve!
	http.ListenAndServe(":8080", nil)
}
