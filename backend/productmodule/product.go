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

func (p *Module) GetProductBatch(limit int, offset int, keyword string) ([]ProductResponse, error) {
	var resps []ProductResponse
	var err error

	resps, err = p.Storage.GetProductBatch(limit, offset, keyword)
	if err != nil {
		log.Println("[ProductModule][GetProductBatch] problem getting storage data, err: ", err.Error())
		return resps, err
	}

	for _, resp := range resps {
		resp.PriceFormat = formatPrice(resp.Price)
	}

	return resps, nil
}

func (p *Module) UpdateProduct(id int64, data UpdateProductRequest) (ProductResponse, error) {
	resp, err := p.Storage.UpdateProduct(id, data)
	if err != nil {
		log.Println("[ProductModule][UpdateProduct] problem getting storage data, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}
