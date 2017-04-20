package data

import (
	"math/rand"

	mgo "gopkg.in/mgo.v2"

	"github.com/graphql-go/graphql"
)

// Schema ...
var Schema graphql.Schema
var queryType *graphql.Object

var QuotesCollection *mgo.Collection

type Collections struct {
	Quotes *mgo.Collection `json:"records"`
}

func init() {
	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"latestPost": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Hello World!", nil
				},
			},
			"quotesCount": &graphql.Field{
				Type:        graphql.Int,
				Description: " Count of the Quotes collection in the Mongo database",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					count, _ := p.Context.Value("dbCollections ").(Collections).Quotes.Count()
					return count, nil
				},
			},
			"randomNumber": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return rand.Intn(100), nil
				},
			},
			"quotesLibrary": &graphql.Field{
				Type:        QuotesLibraryType,
				Description: "The Quotes Library",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					quotes := QuoteList{}
					return quotes, nil
				},
			},
		},
	})

	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
