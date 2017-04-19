package data

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"golang.org/x/net/context"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var quoteIDFetcher = func(obj interface{}, info graphql.ResolveInfo, ctx context.Context) (string, error) {
	switch obj := obj.(type) {
	case *bson.ObjectId:
		fmt.Println(obj.String)
		return fmt.Sprintf("%v", obj.String), nil
	}
	return "", nil
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
