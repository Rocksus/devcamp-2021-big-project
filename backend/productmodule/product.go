package productmodule

import (
	"database/sql"
	"errors"
	"log"
)

type Module struct {
	ProductDB *sql.DB
}

func NewProductModule(db *sql.DB) *Module {
	return &Module{
		ProductDB: db,
	}
}

func (p *Module) AddProduct(data InsertProductRequest) (ProductResponse, error) {
	var resp ProductResponse

	if err := data.Sanitize(); err != nil {
		log.Println("[ProductModule][AddProduct] bad request, err: ", err.Error())
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
		log.Println("[ProductModule][AddProduct] problem querying to db, err: ", err.Error())
		return resp, err
	}

	resp = ProductResponse{
		ID: id,
	}

	return resp, nil
}

func (p *Module) GetProduct(id int64) (ProductResponse, error) {
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
		log.Println("[ProductModule][GetProduct] problem querying to db, err: ", err.Error())
		return resp, err
	}
	resp.ID = id

	return resp, nil
}

func (p *Module) GetProductBatch(lastID int64, limit int) ([]ProductResponse, error) {
	resp := make([]ProductResponse, 0)
	rows, err := p.ProductDB.Query(getProductBatchQuery, lastID, limit)
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

func (p *Module) UpdateProduct(id int64, data UpdateProductRequest) (ProductResponse, error) {
	var resp ProductResponse

	query, values := data.BuildQuery(id)
	res, err := p.ProductDB.Exec(query, values...)
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
