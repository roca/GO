package main

import "fmt"


type GraphQLString string

type GraphQLField struct {
	Type interface{}
	Resolver func(parent, args interface{}) (map[string]interface{}, error)
}
func (g *GraphQLField)Resolve() {
	_,err := g.Resolver(g,"")
	if err != nil {
		fmt.Println(err.Error())
	}
}

type GraphQLObject struct {
	Name string
	Description string
	Fields func() map[string]GraphQLField
}

func main() {
   PostType := GraphQLObject{
	   Name: "Post",
	   Description: "This a Post",
	   Fields: func() map[string]GraphQLField {
		   fields := make(map[string]GraphQLField)
		   fields["id"] = GraphQLField{
			   Type: GraphQLString(""),
		   }
		   return fields
	   },
	}

	fmt.Println(PostType.Name)
}