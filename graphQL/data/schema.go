package data

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"reflect"

	"golang.org/x/net/context"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

// Schema ...
var Schema graphql.Schema
var queryType *graphql.Object

var QuotesCollection *mgo.Collection

type Collections struct {
	Quotes *mgo.Collection `json:"records"`
}

var NodeDefinitions *relay.NodeDefinitions
var QuoteNodeType *graphql.Object
var QuotesLibraryNodeType *graphql.Object

func init() {

	NodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
			// resolve id from global id
			resolvedID := relay.FromGlobalID(id)

			// based on id and its type, return the object
			switch resolvedID.Type {
			case "Quote":
				quote := Quote{}
				err := ctx.Value("dbCollections ").(Collections).Quotes.FindId(bson.ObjectIdHex(resolvedID.ID)).One(&quote)
				if err != nil {
					log.Fatal(err)
				}
				return quote, nil
			default:
				return nil, errors.New("Unknown node type")
			}
		},
		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			// based on the type of the value, return GraphQLObjectType
			fmt.Println(reflect.TypeOf(p.Value))
			switch p.Value.(type) {
			case Quote:
				fmt.Println(p.Value)
				return QuoteNodeType
			default:
				return QuoteNodeType
			}
		},
	})

	QuoteNodeType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Quote",
		Description: "A quote in the library",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Quote", quoteIDFetcher),
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
		},
		Interfaces: []*graphql.Interface{
			NodeDefinitions.NodeInterface,
		},
	})

	quotesConnectionDefinition := relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name:     "Quote",
		NodeType: QuoteNodeType,
	})

	QuotesLibraryNodeType = graphql.NewObject(graphql.ObjectConfig{
		Name: "QuotesLibrary",
		Fields: graphql.Fields{
			"allQuotes": &graphql.Field{
				Type:        quotesConnectionDefinition.ConnectionType,
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
			"node": NodeDefinitions.NodeField,
		},
	})

	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
