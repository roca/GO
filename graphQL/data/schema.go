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

var nodeDefinitions *relay.NodeDefinitions
var QuoteNodeType *graphql.Object
var QuotesLibraryNodeType *graphql.Object

func init() {

	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
			// resolve id from global id
			resolvedID := relay.FromGlobalID(id)
			fmt.Println(resolvedID.Type)
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
			fmt.Println(p.Value)
			switch p.Value.(type) {
			case Quote:
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
			nodeDefinitions.NodeInterface,
		},
	})

	quoteConnectionDefinition := relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name:     "Quote",
		NodeType: QuoteNodeType,
	})

	QuotesLibraryNodeType = graphql.NewObject(graphql.ObjectConfig{
		Name: "QuotesLibrary",
		Fields: graphql.Fields{
			"allQuotes": &graphql.Field{
				Type:        quoteConnectionDefinition.ConnectionType,
				Args:        relay.ConnectionArgs,
				Description: "A list of the quotes in the database",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := relay.NewConnectionArguments(p.Args)
					quotes := []interface{}{}
					if quote, ok := p.Source.(*Quote); ok {
						quotes = append(quotes, *quote)
					}
					return relay.ConnectionFromArray(quotes, args), nil
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
			"node": nodeDefinitions.NodeField,
			// "quotes": &graphql.Field{
			// 	Type: QuotesLibraryNodeType,
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		quotes := QuoteList{}
			// 		return quotes, nil
			// 	},
			// },
		},
	})

	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
