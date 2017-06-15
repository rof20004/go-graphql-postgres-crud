package main

import (
	"strconv"

	"github.com/graphql-go/graphql"
)

// MutationType type
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Description: "New User Email",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email := p.Args["email"].(string)
				user := &User{
					Email: email,
				}
				err := InsertUser(user)
				return user, err
			},
		},
		"removeUser": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "User ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err = RemoveUserByID(id)
				return (err == nil), err
			},
		},
		"createEndereco": &graphql.Field{
			Type: EnderecoType,
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Description: "Id of user creating the new post",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"street": &graphql.ArgumentConfig{
					Description: "New endereco street",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"number": &graphql.ArgumentConfig{
					Description: "New endereco number",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["user"].(string)
				userID, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				street := p.Args["street"].(string)
				number := p.Args["number"].(int)
				endereco := &Endereco{
					UserID: userID,
					Street: street,
					Number: number,
				}
				err = InsertEndereco(endereco)
				return endereco, err
			},
		},
		"removeEndereco": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Endereco ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err = RemoveEnderecoByID(id)
				return (err == nil), err
			},
		},
	},
})
