package product

import (
	"context"

	"github.com/graphql-go/graphql"

	"github.com/Rocksus/devcamp-2021-big-project/backend/productmodule"
	"github.com/Rocksus/devcamp-2021-big-project/backend/tracer"
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
		span, ctx := tracer.StartSpanFromContext(context.Background(), "gqlresolver.addproduct")
		defer span.Finish()

		name, _ := p.Args["name"].(string)
		description, _ := p.Args["description"].(string)
		price, _ := p.Args["price"].(int)
		rating, _ := p.Args["rating"].(float64)
		imageURL, _ := p.Args["imageURL"].(string)
		previewImageURL, _ := p.Args["previewImageURL"].(string)
		slug, _ := p.Args["slug"].(string)

		req := productmodule.InsertProductRequest{
			Name:            name,
			Description:     description,
			Price:           int64(price),
			Rating:          float32(rating),
			ImageURL:        imageURL,
			PreviewImageURL: previewImageURL,
			Slug:            slug,
		}

		return r.p.AddProduct(ctx, req)
	}
}

func (r *Resolver) GetProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		span, ctx := tracer.StartSpanFromContext(context.Background(), "gqlresolver.getproduct")
		defer span.Finish()
		id, _ := p.Args["id"].(int)

		return r.p.GetProduct(ctx, int64(id))
	}
}

func (r *Resolver) GetProductBatch() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		span, ctx := tracer.StartSpanFromContext(context.Background(), "gqlresolver.getproductbatch")
		defer span.Finish()

		lastID, ok := p.Args["lastid"].(int)
		if !ok {
			lastID = 0
		}
		limit, ok := p.Args["limit"].(int)
		if !ok || limit == 0 {
			limit = 10
		}

		return r.p.GetProductBatch(ctx, int64(lastID), limit)
	}
}

func (r *Resolver) UpdateProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		span, ctx := tracer.StartSpanFromContext(context.Background(), "gqlresolver.updateproduct")
		defer span.Finish()

		id, _ := p.Args["id"].(int)
		name, _ := p.Args["name"].(string)
		description, _ := p.Args["description"].(string)
		price, _ := p.Args["price"].(int)
		rating, _ := p.Args["rating"].(float64)
		imageURL, _ := p.Args["imageURL"].(string)
		previewImageURL, _ := p.Args["previewImageURL"].(string)
		slug, _ := p.Args["slug"].(string)

		req := productmodule.UpdateProductRequest{
			Name:            name,
			Description:     description,
			Price:           int64(price),
			Rating:          float32(rating),
			ImageURL:        imageURL,
			PreviewImageURL: previewImageURL,
			Slug:            slug,
		}

		return r.p.UpdateProduct(ctx, int64(id), req)
	}
}
