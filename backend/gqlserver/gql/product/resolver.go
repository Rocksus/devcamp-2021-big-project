package product

import (
	"github.com/graphql-go/graphql"

	"github.com/Rocksus/devcamp-2021-big-project/backend/productmodule"
)

type Resolver struct {
	p *productmodule.Module
}

func NewResolver(p *productmodule.Module) *Resolver {
	return &Resolver{
		p: p,
	}
}

func (r *Resolver) AddProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["product_name"].(string)
		description, _ := p.Args["product_description"].(string)
		price, _ := p.Args["product_price"].(int)
		rating, _ := p.Args["rating"].(float64)
		imageURL, _ := p.Args["product_image"].(string)
		additionalImageURL, _ := p.Args["additional_product_image"].([]string)

		req := productmodule.InsertProductRequest{
			Name:               name,
			Description:        description,
			Price:              int64(price),
			Rating:             float32(rating),
			ImageURL:           imageURL,
			AdditionalImageURL: additionalImageURL,
		}

		return r.p.AddProduct(req)
	}
}

func (r *Resolver) GetProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["productId"].(int)

		return r.p.GetProduct(int64(id))
	}
}

func (r *Resolver) GetProductBatch() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		limit, ok := p.Args["limit"].(int)
		if !ok || limit < 0 {
			limit = 10
		}
		offset, ok := p.Args["offset"].(int)
		if !ok || offset < 0 {
			offset = 0
		}

		return r.p.GetProductBatch(limit, offset)
	}
}

func (r *Resolver) UpdateProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["product_id"].(int)
		name, _ := p.Args["product_name"].(string)
		description, _ := p.Args["product_description"].(string)
		price, _ := p.Args["product_price"].(int)
		rating, _ := p.Args["product_rating"].(float64)
		imageURL, _ := p.Args["product_image"].(string)
		additionalImageURL, _ := p.Args["additional_product_image"].([]string)

		req := productmodule.UpdateProductRequest{
			Name:               name,
			Description:        description,
			Price:              int64(price),
			Rating:             float32(rating),
			ImageURL:           imageURL,
			AdditionalImageURL: additionalImageURL,
		}

		return r.p.UpdateProduct(int64(id), req)
	}
}
