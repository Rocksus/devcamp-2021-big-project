package product

import (
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	p *Product
}

func NewResolver(p *Product) *Resolver {
	return &Resolver{
		p: p,
	}
}

func (r *Resolver) AddProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)
		description, _ := p.Args["description"].(string)
		price, _ := p.Args["price"].(int)
		rating, _ := p.Args["rating"].(float64)
		imageURL, _ := p.Args["imageURL"].(string)
		previewImageURL, _ := p.Args["previewImageURL"].(string)
		slug, _ := p.Args["slug"].(string)

		req := insertProductRequest{
			Name:            name,
			Description:     description,
			Price:           int64(price),
			Rating:          float32(rating),
			ImageURL:        imageURL,
			PreviewImageURL: previewImageURL,
			Slug:            slug,
		}

		return r.p.AddProduct(req)
	}
}

func (r *Resolver) GetProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(int)

		return r.p.GetProduct(int64(id))
	}
}

func (r *Resolver) GetProductBatch() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		lastID, ok := p.Args["lastid"].(int)
		if !ok {
			lastID = 0
		}
		limit, ok := p.Args["limit"].(int)
		if !ok || limit == 0 {
			limit = 10
		}

		return r.p.GetProductBatch(int64(lastID), limit)
	}
}

func (r *Resolver) EditProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(int)
		name, _ := p.Args["name"].(string)
		description, _ := p.Args["description"].(string)
		price, _ := p.Args["price"].(int)
		rating, _ := p.Args["rating"].(float64)
		imageURL, _ := p.Args["imageURL"].(string)
		previewImageURL, _ := p.Args["previewImageURL"].(string)
		slug, _ := p.Args["slug"].(string)

		req := editProductRequest{
			Name:            name,
			Description:     description,
			Price:           int64(price),
			Rating:          float32(rating),
			ImageURL:        imageURL,
			PreviewImageURL: previewImageURL,
			Slug:            slug,
		}

		return r.p.EditProduct(int64(id), req)
	}
}
