package data

import (
	"log"
	"math/rand"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

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
			"quotes": &graphql.Field{
				Type:        graphql.NewList(QuotesLibraryType),
				Description: "List of all Quotes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					quote := Quote{}
					err := p.Context.Value("dbCollections ").(Collections).Quotes.Find(bson.M{"author": "H. Jackson Brown"}).One(&quote)
					if err != nil {
						log.Fatal(err)
					}
					return QuoteList{quote}, nil
				},
			},
		},
	})

	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
