package productmodule

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Rocksus/devcamp-2021-big-project/backend/tracer"
	"github.com/gomodule/redigo/redis"
)

type cache struct {
	ProductCache redis.Conn
}

func newCache(c redis.Conn) *cache {
	return &cache{
		ProductCache: c,
	}
}

func (s *cache) GetProduct(ctx context.Context, id int64) (ProductResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "productmodule.getproduct.cache")
	defer span.Finish()

	var resp ProductResponse

	key := fmt.Sprintf(cacheKeyProduct, id)
	span.SetTag("cachekey", key)

	cachedData, err := redis.Bytes(s.ProductCache.Do("GET", key))
	if err != nil {
		log.Println("[ProductModule][GetProduct][Cache] problem getting cache, err: ", err.Error())
		return resp, err
	}

	if err := json.Unmarshal(cachedData, &resp); err != nil {
		log.Println("[ProductModule][GetProduct][Cache] problem unmarshalling cache, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}

func (s *cache) SetProduct(ctx context.Context, data ProductResponse) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "productmodule.getproduct.cache.set")
	defer span.Finish()
	key := fmt.Sprintf(cacheKeyProduct, data.ID)
	span.SetTag("cachekey", key)

	preparedData, err := json.Marshal(data)
	if err != nil {
		log.Println("[ProductModule][SetProduct][Cache] problem marshalling cache, err: ", err.Error())
		return err
	}
	if _, err := s.ProductCache.Do("SET", key, preparedData); err != nil {
		log.Println("[ProductModule][SetProduct][Cache] problem setting cache, err: ", err.Error())
		return err
	}

	return nil
}

func (s *cache) GetProductBatch(ctx context.Context, lastID int64, limit int) ([]ProductResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "productmodule.getproductbatchdata.cache")
	defer span.Finish()

	var resp []ProductResponse

	key := fmt.Sprintf(cacheKeyProductBatch, lastID, limit)
	span.SetTag("cachekey", key)

	cachedData, err := redis.Bytes(s.ProductCache.Do("GET", key))
	if err != nil {
		log.Println("[ProductModule][GetProductBatch][Cache] problem getting cache, err: ", err.Error())
		return resp, err
	}

	if err := json.Unmarshal(cachedData, &resp); err != nil {
		log.Println("[ProductModule][GetProductBatch][Cache] problem unmarshalling cache, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}

func (s *cache) SetProductBatch(ctx context.Context, lastID int64, limit int, data []ProductResponse) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "productmodule.getproductbatch.cache.set")
	defer span.Finish()

	key := fmt.Sprintf(cacheKeyProductBatch, lastID, limit)
	span.SetTag("cachekey", key)

	preparedData, err := json.Marshal(data)
	if err != nil {
		log.Println("[ProductModule][SetProductBatch][Cache] problem marshalling cache, err: ", err.Error())
		return err
	}
	if _, err := s.ProductCache.Do("SET", key, preparedData); err != nil {
		log.Println("[ProductModule][SetProductBatch][Cache] problem setting cache, err: ", err.Error())
		return err
	}

	return nil
}
