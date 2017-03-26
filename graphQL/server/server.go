package server

import (
	"net/http"

	"github.com/graphql-go/graphql"
)

var Version int

// type Query struct {
// 	latestPost string
// }

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

	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	// and serve!
	http.ListenAndServe(":8080", nil)
}
