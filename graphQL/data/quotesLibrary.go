package data

import (
	"errors"
	"fmt"

	"golang.org/x/net/context"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var quoteIDFetcher = func(obj interface{}, info graphql.ResolveInfo, ctx context.Context) (string, error) {
	switch obj := obj.(type) {
	case (Quote):
		fmt.Printf("ID: %v\n", obj.ID.Hex)
		return fmt.Sprintf("%v", obj.ID.Hex), nil
	}
	return "", errors.New("Not a Quote")
}

var QuotesLibraryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "QuotesLibrary",
	Description: "List of all quotes in the Library",
	Fields: graphql.Fields{
		"id": relay.GlobalIDField("QuotesLibrary", quoteIDFetcher),
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
	},
})
