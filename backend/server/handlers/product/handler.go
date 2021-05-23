package product

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Rocksus/devcamp-2021-big-project/backend/server"
)

type Handler struct {
	ProductDB *sql.DB
}

func NewProductHandler(db *sql.DB) *Handler {
	return &Handler{
		ProductDB: db,
	}
}

func (p *Handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ProductHandler][AddProduct] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	var data insertProductRequest
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("[ProductHandler][AddProduct] unable to unmarshal json, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}
	if err := data.Sanitize(); err != nil {
		log.Println("[ProductHandler][AddProduct] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err)
		return
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
		log.Println("[ProductHandler][AddProduct] problem querying to db, err: ", err.Error())
		server.RenderError(w, http.StatusInternalServerError, err)
		return
	}

	resp := cuProductResponse{
		ID: id,
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

	var resp productResponse
	if err := p.ProductDB.QueryRow(getProductQuery, queryID).Scan(&resp); err != nil {
		log.Println("[ProductHandler][GetProduct] problem querying to db, err: ", err.Error())
		server.RenderError(w, http.StatusInternalServerError, err)
		return
	}

	server.RenderResponse(w, http.StatusCreated, resp)
	return
}

func (p *Handler) GetProductBatch(w http.ResponseWriter, r *http.Request) {
	var lastID, limit int
	var err error
	vars := mux.Vars(r)
	lastID, err = strconv.Atoi(vars["lastid"])
	if err != nil {
		lastID = 0
	}
	limit, err = strconv.Atoi(vars["limit"])
	if err != nil {
		limit = 10
	}

	var resp productResponse
	row, err := p.ProductDB.Query(getProductQuery, lastID, limit)
	if err != nil {
		log.Println("[ProductHandler][GetProductBatch] problem querying to db, err: ", err.Error())
		server.RenderError(w, http.StatusInternalServerError, err)
		return
	}
	defer row.Close()
	if err := row.Scan(&resp); err != nil {
		log.Println("[ProductHandler][GetProductBatch] problem with scanning db row, err: ", err.Error())
		server.RenderError(w, http.StatusInternalServerError, err)
		return
	}

	server.RenderResponse(w, http.StatusCreated, resp)
	return
}

func (p *Handler) EditProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[ProductHandler][EditProduct] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ProductHandler][EditProduct] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	var data editProductRequest
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("[ProductHandler][EditProduct] unable to unmarshal json, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err)
		return
	}

	query, values := data.BuildQuery(queryID)
	res, err := p.ProductDB.Exec(query, values...)
	if err != nil {
		log.Println("[ProductHandler][EditProduct] problem querying to db, err: ", err.Error())
		server.RenderError(w, http.StatusInternalServerError, err)
		return
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("[ProductHandler][EditProduct] problem querying to db, err: ", err.Error())
		server.RenderError(w, http.StatusInternalServerError, err)
		return
	}
	if rowsAffected == 0 {
		log.Println("[ProductHandler][EditProduct] no rows affected in db")
		server.RenderError(w, http.StatusInternalServerError, errors.New("no rows affected in db"))
		return
	}

	resp := cuProductResponse{
		ID: queryID,
	}

	server.RenderResponse(w, http.StatusCreated, resp)
	return
}
