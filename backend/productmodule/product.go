package productmodule

import (
	"context"
	"database/sql"
	"github.com/Rocksus/devcamp-2021-big-project/backend/messaging"
	"log"

	"github.com/Rocksus/devcamp-2021-big-project/backend/tracer"
	"github.com/gomodule/redigo/redis"
)

type Module struct {
	Storage  *storage
	Cache    *cache
	Producer *messaging.Producer
}

func NewProductModule(db *sql.DB, cache redis.Conn, p *messaging.Producer) *Module {
	return &Module{
		Storage:  newStorage(db),
		Cache:    newCache(cache),
		Producer: p,
	}
}

func (p *Module) AddProduct(ctx context.Context, data InsertProductRequest) (ProductResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "productmodule.addproduct")
	defer span.Finish()
	span.SetTag("added-product-name", data.Name)

	if err := data.Sanitize(); err != nil {
		log.Println("[ProductModule][AddProduct] bad request, err: ", err.Error())
		return ProductResponse{}, err
	}

	resp, err := p.Storage.AddProduct(ctx, data)
	if err != nil {
		log.Println("[ProductModule][AddProduct] problem in getting from storage, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}

func (p *Module) GetProduct(ctx context.Context, id int64) (ProductResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "productmodule.getproduct")
	defer span.Finish()
	span.SetTag("id", id)
	var resp ProductResponse
	var err error

	defer func() {
		// publish view data to be handled by consumer
		message := producerMessage{
			Event:         "view",
			ProductDetail: resp,
		}

		if err := p.Producer.Publish(topicProductView, message); err != nil {
			log.Println("[ProductModule][GetProduct] failed to publish message data, err: ", err.Error())
		}
	}()

	resp, err = p.Cache.GetProduct(ctx, id)
	if err == nil {
		return resp, nil
	}
	if err != nil && err != redis.ErrNil {
		log.Println("[ProductModule][GetProduct] problem getting cache data, err: ", err.Error())
	}

	resp, err = p.Storage.GetProduct(ctx, id)
	if err != nil {
		log.Println("[ProductModule][GetProduct] problem getting storage data, err: ", err.Error())
		return resp, err
	}

	if err := p.Cache.SetProduct(ctx, resp); err != nil {
		log.Println("[ProductModule][GetProduct] problem setting cache data, err: ", err.Error())
	}

	return resp, nil
}

func (p *Module) GetProductBatch(ctx context.Context, lastID int64, limit int) ([]ProductResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "productmodule.getproductbatchdata")
	defer span.Finish()

	resp, err := p.Cache.GetProductBatch(ctx, lastID, limit)
	if err == nil {
		return resp, nil
	}
	if err != nil && err != redis.ErrNil {
		log.Println("[ProductModule][GetProductBatch] problem getting cache, err: ", err.Error())
	}

	resp, err = p.Storage.GetProductBatch(ctx, lastID, limit)
	if err != nil {
		log.Println("[ProductModule][GetProductBatch] problem getting storage data, err: ", err.Error())
		return resp, err
	}

	if err := p.Cache.SetProductBatch(ctx, lastID, limit, resp); err != nil {
		log.Println("[ProductModule][GetProductBatch] problem setting cache data, err: ", err.Error())
	}

	return resp, nil
}

func (p *Module) UpdateProduct(ctx context.Context, id int64, data UpdateProductRequest) (ProductResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "productmodule.updateproduct")
	defer span.Finish()
	span.SetTag("id", id)

	resp, err := p.Storage.UpdateProduct(ctx, id, data)
	if err != nil {
		log.Println("[ProductModule][UpdateProduct] problem getting storage data, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}
