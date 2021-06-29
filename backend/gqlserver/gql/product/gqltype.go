package product

import "github.com/graphql-go/graphql"

var ProductType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"product_id": &graphql.Field{
				Type: graphql.Int,
			},
			"product_name": &graphql.Field{
				Type: graphql.String,
			},
			"product_description": &graphql.Field{
				Type: graphql.String,
			},
			"product_price": &graphql.Field{
				Type: graphql.Int,
			},
			"product_price_format": &graphql.Field{
				Type: graphql.String,
			},
			"rating": &graphql.Field{
				Type: graphql.Float,
			},
			"product_image": &graphql.Field{
				Type: graphql.String,
			},
			"additional_product_image": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)
