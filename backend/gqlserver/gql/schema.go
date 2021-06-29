package gql

import (
	"github.com/graphql-go/graphql"

	"github.com/Rocksus/devcamp-2021-big-project/backend/gqlserver/gql/product"
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
				"ProductDetail": &graphql.Field{
					Type:        product.ProductType,
					Description: "Get product by ID",
					Args: graphql.FieldConfigArgument{
						"productId": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: s.productResolver.GetProduct(),
				},
				"ProductLists": &graphql.Field{
					Type:        graphql.NewList(product.ProductType),
					Description: "Get products by pagination",
					Args: graphql.FieldConfigArgument{
						"limit": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"offset": &graphql.ArgumentConfig{
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
				"AddProduct": &graphql.Field{
					Type:        product.ProductType,
					Description: "Create new product",
					Args: graphql.FieldConfigArgument{
						"product_name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"product_description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"product_price": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"rating": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"product_image": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"additional_product_image": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
						},
					},
					Resolve: s.productResolver.AddProduct(),
				},
				"UpdateProduct": &graphql.Field{
					Type:        product.ProductType,
					Description: "Update existing product",
					Args: graphql.FieldConfigArgument{
						"product_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"product_name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"product_description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"product_price": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"rating": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"product_image": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"additional_product_image": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
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
