package productmodule

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Rocksus/devcamp-2021-big-project/backend/cache"
	"github.com/gomodule/redigo/redis"
)

type Cache struct {
	ProductCache *cache.Redis
}

func newCache(r *cache.Redis) *Cache {
	return &Cache{
		ProductCache: r,
	}
}

func (s *Cache) GetProduct(id int64) (ProductResponse, error) {
	var resp ProductResponse

	key := fmt.Sprintf(cacheKeyProduct, id)

	cachedData, err := redis.Bytes(s.ProductCache.Do("GET", key))
	if err != nil {
		return resp, err
	}

	if err := json.Unmarshal(cachedData, &resp); err != nil {
		log.Println("[ProductModule][GetProduct][Cache] problem unmarshalling cache, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}

func (s *Cache) SetProduct(data ProductResponse) error {
	key := fmt.Sprintf(cacheKeyProduct, data.ID)

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

func (s *Cache) GetProductBatch(limit, offset int) ([]ProductResponse, error) {
	var resp []ProductResponse

	key := fmt.Sprintf(cacheKeyProductBatch, limit, offset)

	cachedData, err := redis.Bytes(s.ProductCache.Do("GET", key))
	if err != nil {
		return resp, err
	}

	if err := json.Unmarshal(cachedData, &resp); err != nil {
		log.Println("[ProductModule][GetProductBatch][Cache] problem unmarshalling cache, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}

func (s *Cache) SetProductBatch(limit, offset int, data []ProductResponse) error {
	key := fmt.Sprintf(cacheKeyProductBatch, limit, offset)

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

func (s *Cache) DelProductCache(id int64) error {
	key := fmt.Sprintf(cacheKeyProduct, id)

	if _, err := s.ProductCache.Do("DEL", key); err != nil {
		log.Println("[ProductModule][SetProduct][Cache] problem deleting cache, err: ", err.Error())
		return err
	}

	return nil
}
