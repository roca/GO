package data

import "github.com/graphql-go/graphql"

var QuotesLibraryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "QuotesLibrary",
	Description: "List of all quotes in the Library",
	Fields: graphql.Fields{
		//"id": relay.GlobalIDField("Author", nil),
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
	},
})
