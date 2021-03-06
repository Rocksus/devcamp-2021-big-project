package product

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/Rocksus/devcamp-2021-big-project/backend/productmodule"
	"github.com/Rocksus/devcamp-2021-big-project/backend/server"
	"github.com/Rocksus/devcamp-2021-big-project/backend/tracer"
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
	span, ctx := tracer.StartSpanFromContext(r.Context(), "producthandler.addproduct")
	defer span.Finish()
	timeStart := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ProductHandler][AddProduct] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	var data productmodule.InsertProductRequest
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("[ProductHandler][AddProduct] unable to unmarshal json, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	res, err := p.product.AddProduct(ctx, data)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	resp := cuProductResponse{
		ID: res.ID,
	}

	server.RenderResponse(w, http.StatusCreated, resp, timeStart)
	return
}

func (p *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	span, ctx := tracer.StartSpanFromContext(r.Context(), "producthandler.getproduct")
	defer span.Finish()
	timeStart := time.Now()

	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[ProductHandler][GetProduct] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	resp, err := p.product.GetProduct(ctx, queryID)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusOK, resp, timeStart)
	return
}

func (p *Handler) GetProductBatch(w http.ResponseWriter, r *http.Request) {
	span, ctx := tracer.StartSpanFromContext(r.Context(), "producthandler.getproductbatch")
	defer span.Finish()

	var limit int
	var offset int
	timeStart := time.Now()

	var err error
	// query parameters are not available in mux vars
	vars := r.URL.Query()
	limit, err = strconv.Atoi(vars.Get("limit"))
	if err != nil || limit < 0 {
		limit = 10
	}
	offset, err = strconv.Atoi(vars.Get("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	resp, err := p.product.GetProductBatch(ctx, limit, offset)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusOK, resp, timeStart)
	return
}

func (p *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	span, ctx := tracer.StartSpanFromContext(r.Context(), "producthandler.updateproduct")
	defer span.Finish()

	timeStart := time.Now()
	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[ProductHandler][UpdateProduct] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ProductHandler][UpdateProduct] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	var data productmodule.UpdateProductRequest
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("[ProductHandler][UpdateProduct] unable to unmarshal json, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	resp, err := p.product.UpdateProduct(ctx, queryID, data)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusCreated, resp, timeStart)
	return
}
