package main

import (
	"strconv"

	"github.com/graphql-go/graphql"
)

// EnderecoType type
var EnderecoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Endereco",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if endereco, ok := p.Source.(*Endereco); ok == true {
					return endereco.ID, nil
				}
				return nil, nil
			},
		},
		"user_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if endereco, ok := p.Source.(*Endereco); ok == true {
					return endereco.UserID, nil
				}
				return nil, nil
			},
		},
		"street": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if endereco, ok := p.Source.(*Endereco); ok == true {
					return endereco.Street, nil
				}
				return nil, nil
			},
		},
		"number": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if endereco, ok := p.Source.(*Endereco); ok == true {
					return endereco.Number, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	EnderecoType.AddFieldConfig("endereco", &graphql.Field{
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
