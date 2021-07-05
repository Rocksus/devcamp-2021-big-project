package productmodule

import (
	"database/sql"
	"log"

	"github.com/Rocksus/devcamp-2021-big-project/backend/cache"
	"github.com/gomodule/redigo/redis"
)

type Module struct {
	Storage *storage
	Cache   *Cache
}

func NewProductModule(db *sql.DB, redisCache *cache.Redis) *Module {
	return &Module{
		Storage: newStorage(db),
		Cache:   newCache(redisCache),
	}
}

func (p *Module) AddProduct(data InsertProductRequest) (ProductResponse, error) {
	if err := data.Sanitize(); err != nil {
		log.Println("[ProductModule][AddProduct] bad request, err: ", err.Error())
		return ProductResponse{}, err
	}

	resp, err := p.Storage.AddProduct(data)
	if err != nil {
		log.Println("[ProductModule][AddProduct] problem in getting from storage, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}

func (p *Module) GetProduct(id int64) (ProductResponse, error) {
	var resp ProductResponse
	var err error

	resp, err = p.Cache.GetProduct(id)
	if err == nil {
		return resp, nil
	} else if err != redis.ErrNil {
		log.Println("[ProductModule][GetProduct] problem getting cache data, err: ", err.Error())
	}

	resp, err = p.Storage.GetProduct(id)
	if err != nil {
		log.Println("[ProductModule][GetProduct] problem getting storage data, err: ", err.Error())
		return resp, err
	}
	resp.PriceFormat = formatPrice(resp.Price)

	if err := p.Cache.SetProduct(resp); err != nil {
		log.Println("[ProductModule][GetProduct] problem setting cache data, err: ", err.Error())
	}

	return resp, nil
}

func (p *Module) GetProductBatch(limit, offset int) ([]ProductResponse, error) {
	resp, err := p.Cache.GetProductBatch(limit, offset)
	if err == nil {
		return resp, nil
	} else if err != redis.ErrNil {
		log.Println("[ProductModule][GetProductBatch] problem getting cache, err: ", err.Error())
	}

	resp, err = p.Storage.GetProductBatch(limit, offset)
	if err != nil {
		log.Println("[ProductModule][GetProductBatch] problem getting storage data, err: ", err.Error())
		return resp, err
	}

	if err := p.Cache.SetProductBatch(limit, offset, resp); err != nil {
		log.Println("[ProductModule][GetProductBatch] problem setting cache data, err: ", err.Error())
	}

	return resp, nil
}

func (p *Module) UpdateProduct(id int64, data UpdateProductRequest) (ProductResponse, error) {
	resp, err := p.Storage.UpdateProduct(id, data)
	if err != nil {
		log.Println("[ProductModule][UpdateProduct] problem getting storage data, err: ", err.Error())
		return resp, err
	}
	if err := p.Cache.DelProductCache(id); err != nil {
		log.Println("[ProductModule][UpdateProduct] problem deleting cache data, err: ", err.Error())
	}

	return resp, nil
}
