package data

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/graphql-go/graphql"
)

var quoteIDFetcher = func(obj interface{}, info graphql.ResolveInfo, ctx context.Context) (string, error) {
	switch obj := obj.(type) {
	case *Quote:
		fmt.Println(obj.ID)
		return fmt.Sprintf("%v", obj.ID), nil
	}
	return "", nil
}

var QuotesLibraryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "QuotesLibrary",
	Description: "List of all quotes in the Library",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				//ID := p.Source.(*Quote).ID
				//fmt.Println(ID)
				return "a", nil
			},
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
	},
})
