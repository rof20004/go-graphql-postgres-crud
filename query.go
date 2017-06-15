package main

import (
	"strconv"

	"github.com/graphql-go/graphql"
)

// QueryType tyoe
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "User ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetUserByID(id)
			},
		},
		"users": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(UserType))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return GetUsers()
			},
		},
	},
})
