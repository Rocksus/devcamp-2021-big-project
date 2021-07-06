package productmodule

import (
	"errors"
	"fmt"

	"github.com/lib/pq"
)

const (
	cacheKeyProduct      = "product:%d"
	cacheKeyProductBatch = "products:%d:%d"
)

const (
	topicProductView = "product_view"
)

type ProductResponse struct {
	ID                 int64    `json:"product_id,omitempty" db:"id"`
	Name               string   `json:"product_name,omitempty" db:"name"`
	Description        string   `json:"product_description,omitempty" db:"description"`
	Price              int64    `json:"product_price,omitempty" db:"price"`
	PriceFormat        string   `json:"product_price_format,omitempty" db:"-"`
	Rating             float32  `json:"rating,omitempty" db:"rating"`
	ImageURL           string   `json:"product_image,omitempty" db:"image_url"`
	AdditionalImageURL []string `json:"additional_product_image,omitempty" db:"preview_image_url"`
}

type InsertProductRequest struct {
	Name               string   `json:"product_name"`
	Description        string   `json:"product_description"`
	Price              int64    `json:"product_price"`
	Rating             float32  `json:"rating"`
	ImageURL           string   `json:"product_image"`
	AdditionalImageURL []string `json:"additional_product_image"`
}

func (p InsertProductRequest) Sanitize() error {
	if p.Name == "" {
		return errors.New("name cannot be empty")
	}
	if p.Price == 0 {
		return errors.New("price cannot be empty")
	}
	if p.Rating < 0 || p.Rating > 5 {
		return errors.New("invalid rating range")
	}
	return nil
}

type UpdateProductRequest struct {
	Name               string   `json:"product_name"`
	Description        string   `json:"product_description"`
	Price              int64    `json:"product_price"`
	Rating             float32  `json:"rating"`
	ImageURL           string   `json:"product_image"`
	AdditionalImageURL []string `json:"additional_product_image"`
}

type GetProductBatchRequest struct {
	Name        string `json:"product_name"`
	Description string `json:"product_description"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
}

func (p UpdateProductRequest) BuildQuery(id int64) (string, []interface{}) {
	var fieldQuery string
	fieldValues := make([]interface{}, 0)

	var i = 1
	if p.Name != "" {
		fieldQuery += fmt.Sprintf("name=$%d,", i)
		fieldValues = append(fieldValues, p.Name)
		i++
	}
	if p.Description != "" {
		fieldQuery += fmt.Sprintf("description=$%d,", i)
		fieldValues = append(fieldValues, p.Description)
		i++
	}
	if p.Price != 0 {
		fieldQuery += fmt.Sprintf("price=$%d,", i)
		fieldValues = append(fieldValues, p.Price)
		i++
	}
	if p.Rating != 0 {
		fieldQuery += fmt.Sprintf("rating=$%d,", i)
		fieldValues = append(fieldValues, p.Rating)
		i++
	}
	if p.ImageURL != "" {
		fieldQuery += fmt.Sprintf("image_url=$%d,", i)
		fieldValues = append(fieldValues, p.ImageURL)
		i++
	}
	if len(p.AdditionalImageURL) > 0 {
		fieldQuery += fmt.Sprintf("additional_image_url=$%d,", i)
		fieldValues = append(fieldValues, pq.Array(p.AdditionalImageURL))
		i++
	}

	finalQuery := fmt.Sprintf(updateProductQuery, fieldQuery[:len(fieldQuery)-1], id)

	return finalQuery, fieldValues
}

type producerMessage struct {
	Event         string          `json:"event"`
	ProductDetail ProductResponse `json:"product_detail"`
}
