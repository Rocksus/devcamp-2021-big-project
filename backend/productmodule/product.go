package productmodule

import (
	"database/sql"
	"log"
)

type Module struct {
	Storage *storage
}

func NewProductModule(db *sql.DB) *Module {
	return &Module{
		Storage: newStorage(db),
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

	resp, err = p.Storage.GetProduct(id)
	if err != nil {
		log.Println("[ProductModule][GetProduct] problem getting storage data, err: ", err.Error())
		return resp, err
	}
	resp.PriceFormat = formatPrice(resp.Price)

	return resp, nil
}

func (p *Module) GetProductBatch(keyword string, limit int, offset int) ([]ProductResponse, error) {
	var resp []ProductResponse
	var err error

	resp, err = p.Storage.GetProductBatch(keyword, limit, offset)
	if err != nil {
		log.Println("[ProductModel][GetProductBatch] problem getting storage data, err: ", err.Error())
		return resp, err
	}
	for idx, _ := range resp {
		resp[idx].PriceFormat = formatPrice(resp[idx].Price)
	}

	return resp, nil
}

func (p *Module) UpdateProduct(id int64, data UpdateProductRequest) (ProductResponse, error) {
	resp, err := p.Storage.UpdateProduct(id, data)
	if err != nil {
		log.Println("[ProductModule][UpdateProduct] problem getting storage data, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}
