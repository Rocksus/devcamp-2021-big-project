package productmodule

import (
	"errors"
	"fmt"
)

const (
	cacheKeyProduct      = "product:%d"
	cacheKeyProductBatch = "products:%d:%d"
)

const (
	topicProductView = "product_view"
)

type ProductResponse struct {
	ID              int64   `json:"id,omitempty" db:"id"`
	Name            string  `json:"name,omitempty" db:"name"`
	Description     string  `json:"description,omitempty" db:"description"`
	Price           int64   `json:"price,omitempty" db:"price"`
	Rating          float32 `json:"rating,omitempty" db:"rating"`
	ImageURL        string  `json:"image_url,omitempty" db:"image_url"`
	PreviewImageURL string  `json:"preview_image_url,omitempty" db:"preview_image_url"`
	Slug            string  `json:"slug,omitempty" db:"slug"`
}

type InsertProductRequest struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           int64   `json:"price"`
	Rating          float32 `json:"rating"`
	ImageURL        string  `json:"image_url"`
	PreviewImageURL string  `json:"preview_image_url"`
	Slug            string  `json:"slug"`
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
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           int64   `json:"price"`
	Rating          float32 `json:"rating"`
	ImageURL        string  `json:"image_url"`
	PreviewImageURL string  `json:"preview_image_url"`
	Slug            string  `json:"slug"`
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
	if p.PreviewImageURL != "" {
		fieldQuery += fmt.Sprintf("preview_image_url=$%d,", i)
		fieldValues = append(fieldValues, p.PreviewImageURL)
		i++
	}
	if p.Slug != "" {
		fieldQuery += fmt.Sprintf("slug=$%d,", i)
		fieldValues = append(fieldValues, p.Slug)
		i++
	}

	finalQuery := fmt.Sprintf(updateProductQuery, fieldQuery[:len(fieldQuery)-1], id)

	return finalQuery, fieldValues
}

type producerMessage struct {
	Event         string          `json:"event"`
	ProductDetail ProductResponse `json:"product_detail"`
}
