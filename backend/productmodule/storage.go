package productmodule

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Rocksus/devcamp-2021-big-project/backend/tracer"
	"log"
)

type storage struct {
	ProductDB *sql.DB
}

func newStorage(db *sql.DB) *storage {
	return &storage{
		ProductDB: db,
	}
}

func (s *storage) AddProduct(ctx context.Context, data InsertProductRequest) (ProductResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "productmodule.addproduct.storage")
	defer span.Finish()

	var resp ProductResponse

	var id int64
	if err := s.ProductDB.QueryRowContext(ctx, addProductQuery,
		data.Name,
		data.Description,
		data.Price,
		data.Rating,
		data.ImageURL,
		data.PreviewImageURL,
		data.Slug,
	).Scan(&id); err != nil {
		log.Println("[ProductModule][AddProduct][Storage] problem querying to db, err: ", err.Error())
		return resp, err
	}

	resp = ProductResponse{
		ID: id,
	}
	return resp, nil
}

func (s *storage) GetProduct(ctx context.Context, id int64) (ProductResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "productmodule.getproduct.storage")
	defer span.Finish()

	var resp ProductResponse

	if err := s.ProductDB.QueryRowContext(ctx, getProductQuery, id).Scan(
		&resp.Name,
		&resp.Description,
		&resp.Price,
		&resp.Rating,
		&resp.ImageURL,
		&resp.PreviewImageURL,
		&resp.Slug,
	); err != nil {
		log.Println("[ProductModule][GetProduct] problem querying to db, err: ", err.Error())
		return resp, err
	}
	resp.ID = id

	return resp, nil
}

func (s *storage) GetProductBatch(ctx context.Context, lastID int64, limit int) ([]ProductResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "productmodule.getproductbatch.storage")
	defer span.Finish()

	resp := make([]ProductResponse, 0)

	rows, err := s.ProductDB.QueryContext(ctx, getProductBatchQuery, lastID, limit)
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
			&rowData.PreviewImageURL,
			&rowData.Slug,
		); err != nil {
			log.Println("[ProductModule][GetProductBatch] problem with scanning db row, err: ", err.Error())
			return resp, err
		}
		resp = append(resp, rowData)
	}

	return resp, nil
}

func (s *storage) UpdateProduct(ctx context.Context, id int64, data UpdateProductRequest) (ProductResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "productmodule.updateproduct.storage")
	defer span.Finish()

	var resp ProductResponse

	query, values := data.BuildQuery(id)
	res, err := s.ProductDB.ExecContext(ctx, query, values...)
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
