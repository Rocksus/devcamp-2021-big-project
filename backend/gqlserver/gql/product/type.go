package product

import (
	"errors"
	"fmt"
)

type ProductResponse struct {
	ID              int64   `db:"id"`
	Name            string  `db:"name"`
	Description     string  `db:"description"`
	Price           int64   `db:"price"`
	Rating          float32 `db:"rating"`
	ImageURL        string  `db:"image_url"`
	PreviewImageURL string  `db:"preview_image_url"`
	Slug            string  `db:"slug"`
}

type insertProductRequest struct {
	Name            string
	Description     string
	Price           int64
	Rating          float32
	ImageURL        string
	PreviewImageURL string
	Slug            string
}

func (p insertProductRequest) Sanitize() error {
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

type editProductRequest struct {
	Name            string
	Description     string
	Price           int64
	Rating          float32
	ImageURL        string
	PreviewImageURL string
	Slug            string
}

func (p editProductRequest) BuildQuery(id int64) (string, []interface{}) {
	var fieldQuery string
	fieldValues := make([]interface{}, 0)

	var i int
	if p.Name != "" {
		fieldQuery += fmt.Sprint("name=$", i)
		fieldValues = append(fieldValues, p.Name)
		i++
	}
	if p.Description != "" {
		fieldQuery += fmt.Sprint("description=$", i)
		fieldValues = append(fieldValues, p.Description)
		i++
	}
	if p.Price != 0 {
		fieldQuery += fmt.Sprint("price=$", i)
		fieldValues = append(fieldValues, p.Price)
		i++
	}
	if p.Rating != 0 {
		fieldQuery += fmt.Sprint("rating=$", i)
		fieldValues = append(fieldValues, p.Rating)
		i++
	}
	if p.ImageURL != "" {
		fieldQuery += fmt.Sprint("image_url=$", i)
		fieldValues = append(fieldValues, p.ImageURL)
		i++
	}
	if p.PreviewImageURL != "" {
		fieldQuery += fmt.Sprint("preview_image_url=$", i)
		fieldValues = append(fieldValues, p.PreviewImageURL)
		i++
	}
	if p.Slug != "" {
		fieldQuery += fmt.Sprint("preview_image_url=$", i)
		fieldValues = append(fieldValues, p.Slug)
		i++
	}

	finalQuery := fmt.Sprintf(editProductQuery, fieldQuery, id)

	return finalQuery, fieldValues
}
