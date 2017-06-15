package main

import (
	"strconv"

	"github.com/graphql-go/graphql"
)

// UserType type
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*User); ok == true {
					return user.ID, nil
				}
				return nil, nil
			},
		},
		"email": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*User); ok == true {
					return user.Email, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	UserType.AddFieldConfig("endereco", &graphql.Field{
		Type: EnderecoType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Endereco ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetEnderecoByIDAndUser(id, user.ID)
			}
			return nil, nil
		},
	})
	UserType.AddFieldConfig("enderecos", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(EnderecoType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				return GetEnderecosForUser(user.ID)
			}
			return []Endereco{}, nil
		},
	})
}
