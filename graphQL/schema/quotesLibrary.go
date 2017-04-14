package schema

import "github.com/graphql-go/graphql"

var QuotesLibraryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "QuotesLibrary",
	Description: "List of all quotes in the Library",
	Fields: graphql.Fields{
		"quotes": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		},
	},
})
