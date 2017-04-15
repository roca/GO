package server

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"

	"github.com/GOCODE/graphQL/handlers"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"golang.org/x/net/context"
)

var Version int

var queryType = graphql.NewObject(graphql.ObjectConfig{
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
				count, _ := p.Context.Value("dbCollections ").(DbCollections).quotes.Count()
				return count, nil
			},
		},
		"randomNumber": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return rand.Intn(100), nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

var dbQuotesCollection *mgo.Collection

type DbCollections struct {
	quotes *mgo.Collection `json:"records"`
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {

	opts := handler.NewRequestOptions(r)

	dbCollections := DbCollections{quotes: dbQuotesCollection}

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: opts.Query,
		Context:       context.WithValue(context.Background(), "dbCollections ", dbCollections),
	})
	if len(result.Errors) > 0 {
		log.Printf("wrong result, unexpected errors: %v", result.Errors)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func Start() {

	session, err := mgo.Dial("ec2-52-91-31-26.compute-1.amazonaws.com")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	dbQuotesCollection = session.DB("test").C("quotes")
	count, _ := dbQuotesCollection.Count()
	fmt.Printf("\n\nSuccessfUlly connected to quotes collection. You have %d record(s)\n\n", count)

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/graphiql", handlers.IndexPage).Methods("GET")

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	// h := handler.New(&handler.Config{
	// 	Schema: &Schema,
	// 	Pretty: true,
	// })

	// serve a GraphQL endpoint at `/graphql`

	//http.Handle("/graphql", h)
	http.HandleFunc("/graphql", graphqlHandler)
	http.Handle("/graphiql", gorillaRoute)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	port := os.Getenv("PORT")

	// and serve!
	http.ListenAndServe(":"+port, nil)
}

// curl -XPOST http://local.rit.aws.regeneron.com:8080/graphql -H 'Content-Type: application/graphql' -d 'query Root{ latestPost }'
