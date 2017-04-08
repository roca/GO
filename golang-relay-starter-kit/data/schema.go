package data

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var postType *graphql.Object
var queryType *graphql.Object

// Schema ...
var Schema graphql.Schema

func init() {
	postType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			// Define `id` field as a Relay GlobalID field.
			// It helps with translating your GraphQL object's id into a global id
			// For eg:
			//  For a `Post` type, with an id of `1`, it's global id will be `UG9zdDox`
			//  which is a base64 encoded version of `Post:1` string
			// We will explore more in the next part of this series.
			"id": relay.GlobalIDField("Post", nil),
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"latestPost": &graphql.Field{
				Type: postType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					lastPost := GetLatestPost()
					fmt.Println(lastPost)
					return lastPost, nil
				},
			},
		},
	})

	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

}
