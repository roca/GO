package server

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"

	"github.com/GOCODE/graphQL/handlers"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
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

func Start() {

	session, err := mgo.Dial("ec2-52-91-31-26.compute-1.amazonaws.com")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("test").C("quotes")
	count, _ := db.Count()
	fmt.Printf("\n\nConnect to quotes collection you have %d record(s)\n\n", count)

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/graphiql", handlers.IndexPage).Methods("GET")

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	// serve a GraphQL endpoint at `/graphql`

	http.Handle("/graphql", h)
	http.Handle("/graphiql", gorillaRoute)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	port := os.Getenv("PORT")

	// and serve!
	http.ListenAndServe(":"+port, nil)
}

// curl -XPOST http://local.rit.aws.regeneron.com:8080/graphql -H 'Content-Type: application/graphql' -d 'query Root{ latestPost }'
