package product

import (
	"database/sql"
	"errors"
	"log"
)

type Product struct {
	ProductDB *sql.DB
}

func NewProductService(db *sql.DB) *Product {
	return &Product{
		ProductDB: db,
	}
}

func (p *Product) AddProduct(data insertProductRequest) (ProductResponse, error) {
	var resp ProductResponse

	if err := data.Sanitize(); err != nil {
		log.Println("[ProductGQL][AddProduct] bad request, err: ", err.Error())
		return resp, err
	}

	var id int64
	if err := p.ProductDB.QueryRow(addProductQuery,
		data.Name,
		data.Description,
		data.Price,
		data.Rating,
		data.ImageURL,
		data.PreviewImageURL,
		data.Slug,
	).Scan(&id); err != nil {
		log.Println("[ProductGQL][AddProduct] problem querying to db, err: ", err.Error())
		return resp, err
	}

	resp = ProductResponse{
		ID:              id,
		Name:            data.Name,
		Description:     data.Description,
		Price:           data.Price,
		Rating:          data.Rating,
		ImageURL:        data.ImageURL,
		PreviewImageURL: data.PreviewImageURL,
		Slug:            data.Slug,
	}

	return resp, nil
}

func (p *Product) GetProduct(id int64) (ProductResponse, error) {
	var resp ProductResponse
	if err := p.ProductDB.QueryRow(getProductQuery, id).Scan(
		&resp.Name,
		&resp.Description,
		&resp.Price,
		&resp.Rating,
		&resp.ImageURL,
		&resp.PreviewImageURL,
		&resp.Slug,
	); err != nil {
		log.Println("[ProductGQL][GetProduct] problem querying to db, err: ", err.Error())
		return resp, err
	}
	resp.ID = id

	return resp, nil
}

func (p *Product) GetProductBatch(lastID int64, limit int) ([]ProductResponse, error) {
	resp := make([]ProductResponse, 0)
	rows, err := p.ProductDB.Query(getProductBatchQuery, lastID, limit)
	if err != nil {
		log.Println("[ProductGQL][GetProductBatch] problem querying to db, err: ", err.Error())
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var rowData ProductResponse
		if err := rows.Scan(
			&rowData.Name,
			&rowData.Description,
			&rowData.Price,
			&rowData.Rating,
			&rowData.ImageURL,
			&rowData.PreviewImageURL,
			&rowData.Slug,
		); err != nil {
			log.Println("[ProductGQL][GetProductBatch] problem with scanning db row, err: ", err.Error())
			return resp, err
		}
		resp = append(resp, rowData)
	}

	return resp, nil
}

func (p *Product) EditProduct(id int64, data editProductRequest) (ProductResponse, error) {
	var resp ProductResponse

	query, values := data.BuildQuery(id)
	res, err := p.ProductDB.Exec(query, values...)
	if err != nil {
		log.Println("[ProductGQL][EditProduct] problem querying to db, err: ", err.Error())
		return resp, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("[ProductGQL][EditProduct] problem querying to db, err: ", err.Error())
		return resp, err
	}
	if rowsAffected == 0 {
		log.Println("[ProductGQL][EditProduct] no rows affected in db")
		return resp, errors.New("no rows affected in db")
	}

	resp, err = p.GetProduct(id)
	if err != nil {
		log.Println("[ProductGQL][EditProduct] can't get updated data")
		return resp, err
	}

	return resp, nil
}
