package productmodule

import (
	"database/sql"
	"errors"
	"log"

	"github.com/lib/pq"
)

type storage struct {
	ProductDB *sql.DB
}

func newStorage(db *sql.DB) *storage {
	return &storage{
		ProductDB: db,
	}
}

func (s *storage) AddProduct(data InsertProductRequest) (ProductResponse, error) {
	var resp ProductResponse

	var id int64
	if err := s.ProductDB.QueryRow(addProductQuery,
		data.Name,
		data.Description,
		data.Price,
		data.Rating,
		data.ImageURL,
		pq.Array(&data.AdditionalImageURL),
	).Scan(&id); err != nil {
		log.Println("[ProductModule][AddProduct][Storage] problem querying to db, err: ", err.Error())
		return resp, err
	}

	resp = ProductResponse{
		ID: id,
	}
	return resp, nil
}

func (s *storage) GetProduct(id int64) (ProductResponse, error) {
	var resp ProductResponse

	if err := s.ProductDB.QueryRow(getProductQuery, id).Scan(
		&resp.Name,
		&resp.Description,
		&resp.Price,
		&resp.Rating,
		&resp.ImageURL,
		pq.Array(&resp.AdditionalImageURL),
	); err != nil {
		log.Println("[ProductModule][GetProduct] problem querying to db, err: ", err.Error())
		return resp, err
	}
	resp.ID = id

	return resp, nil
}

func (s *storage) GetProductBatch(limit int, offset int, keyword string) ([]ProductResponse, error) {
	resp := make([]ProductResponse, 0)

	rows, err := s.ProductDB.Query(getProductBatchQuery, limit, offset, keyword)
	if err != nil {
		log.Println("[ProductModule][GetProductBatch] problem querying to db, err: ", err.Error())
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var rowData ProductResponse
		if err := rows.Scan(
			&rowData.ID,
			&rowData.Name,
			&rowData.Description,
			&rowData.Price,
			&rowData.Rating,
			&rowData.ImageURL,
			pq.Array(&rowData.AdditionalImageURL),
		); err != nil {
			log.Println("[ProductModule][GetProductBatch] problem with scanning db row, err: ", err.Error())
			return resp, err
		}
		resp = append(resp, rowData)
	}

	return resp, nil
}

func (s *storage) UpdateProduct(id int64, data UpdateProductRequest) (ProductResponse, error) {
	var resp ProductResponse

	query, values := data.BuildQuery(id)
	res, err := s.ProductDB.Exec(query, values...)
	if err != nil {
		log.Println("[ProductModule][UpdateProduct] problem querying to db, err: ", err.Error())
		return resp, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("[ProductModule][UpdateProduct] problem querying to db, err: ", err.Error())
		return resp, err
	}
	if rowsAffected == 0 {
		log.Println("[ProductModule][UpdateProduct] no rows affected in db")
		return resp, errors.New("no rows affected in db")
	}

	return ProductResponse{
		ID: id,
	}, nil
}
