package product

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Rocksus/devcamp-2021-big-project/backend/productmodule"
	"github.com/Rocksus/devcamp-2021-big-project/backend/server"
)

type Handler struct {
	product *productmodule.Module
}

func NewProductHandler(p *productmodule.Module) *Handler {
	return &Handler{
		product: p,
	}
}

func (p *Handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ProductHandler][AddProduct] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	var data productmodule.InsertProductRequest
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("[ProductHandler][AddProduct] unable to unmarshal json, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	res, err := p.product.AddProduct(data)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	resp := cuProductResponse{
		ID: res.ID,
	}

	server.RenderResponse(w, http.StatusCreated, resp)
	return
}

func (p *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[ProductHandler][GetProduct] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := p.product.GetProduct(queryID)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	server.RenderResponse(w, http.StatusCreated, resp)
	return
}

func (p *Handler) GetProductBatch(w http.ResponseWriter, r *http.Request) {
	var limit int
	var lastID int64

	var err error
	vars := mux.Vars(r)
	lastID, err = strconv.ParseInt(vars["lastid"], 10, 64)
	if err != nil {
		lastID = 0
	}
	limit, err = strconv.Atoi(vars["limit"])
	if err != nil {
		limit = 10
	}

	resp, err := p.product.GetProductBatch(lastID, limit)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	server.RenderResponse(w, http.StatusCreated, resp)
	return
}

func (p *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[ProductHandler][UpdateProduct] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ProductHandler][UpdateProduct] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	var data productmodule.UpdateProductRequest
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("[ProductHandler][UpdateProduct] unable to unmarshal json, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := p.product.UpdateProduct(queryID, data)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	server.RenderResponse(w, http.StatusCreated, resp)
	return
}
