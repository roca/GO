package model

import "errors"

type iGraphQLField interface {
	Resolve(parent, arg interface{}) (map[string]interface{}, error)
}

type iGraphQLObject interface {
	Fields() map[string]iGraphQLField
}

type GraphQLString struct {
	Type string
}

func (s *GraphQLString) Resolve(parent, arg interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	data["id"] = s.Type
	return data, nil
}

type GraphQLList struct {
	Type interface{}
}

func (g *GraphQLList) Resolve(parent, arg interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	return data, errors.New("Not yet implemented")
}

type Post struct {
	ID      string `json:"id"`
	Comment string `json:"comment"`
	UserID  string `json:"user"`
}

func (p *Post) Fields() map[string]iGraphQLField {
	fields := make(map[string]iGraphQLField)

	fields["ID"] = &GraphQLString{Type: "ID"}
	fields["Comment"] = &GraphQLString{Type: "Comment"}
	fields["UserID"] = &GraphQLString{Type: "UserID"}

	return fields
}
