package gql

import (
	"github.com/Rocksus/devcamp-2021-big-project/backend/gqlserver/gql/product"
	"github.com/graphql-go/graphql"
)

type SchemaWrapper struct {
	productResolver *product.Resolver
	Schema          graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) WithProductResolver(pr *product.Resolver) *SchemaWrapper {
	s.productResolver = pr

	return s
}

func (s *SchemaWrapper) Init() error {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"product": &graphql.Field{
					Type:        product.ProductType,
					Description: "Get product by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: s.productResolver.GetProduct(),
				},
				"products": &graphql.Field{
					Type:        graphql.NewList(product.ProductType),
					Description: "Get products by pagination",
					Args: graphql.FieldConfigArgument{
						"lastid": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"limit": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: s.productResolver.GetProductBatch(),
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"addProduct": &graphql.Field{
					Type:        product.ProductType,
					Description: "Create new product",
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"price": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"rating": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"imageURL": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"previewImageURL": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"slug": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: s.productResolver.AddProduct(),
				},
				"updateProduct": &graphql.Field{
					Type:        product.ProductType,
					Description: "Update existing product",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"price": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"rating": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"imageURL": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"previewImageURL": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"slug": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: s.productResolver.UpdateProduct(),
				},
			},
		}),
	})

	if err != nil {
		return err
	}

	s.Schema = schema

	return nil
}
