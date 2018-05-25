package data

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"golang.org/x/net/context"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var quoteIDFetcher = func(obj interface{}, info graphql.ResolveInfo, ctx context.Context) (string, error) {
	fmt.Printf("****quoteIDFetcher:  %v\n", reflect.TypeOf(obj))
	switch obj := obj.(type) {
	case (Quote):
		fmt.Printf("ID: %v\n\n", obj.ID.Hex())
		return fmt.Sprintf("%v", obj.ID.Hex()), nil
	}
	return "", errors.New("Not a Quote")
}

var QuoteType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Quote",
	Description: "A quote in the library",
	Fields: graphql.Fields{
		"id": relay.GlobalIDField("Quote", quoteIDFetcher),
		// "id": &graphql.Field{
		// 	Type: graphql.String,
		// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// 		return p.Source.(Quote).ID.Hex(), nil
		// 	},
		// },
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
	},
	Interfaces: []*graphql.Interface{
		nodeDefinitions.NodeInterface,
	},
})

var QuotesLibraryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "QuotesLibrary",
	Fields: graphql.Fields{
		"allQuotes": &graphql.Field{
			Type:        graphql.NewList(QuoteType),
			Description: "A list of the quotes in the database",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				quotes := QuoteList{}
				err := p.Context.Value("dbCollections ").(Collections).Quotes.Find(nil).All(&quotes)
				if err != nil {
					log.Fatal(err)
				}
				return quotes, nil
			},
		},
	},
})
