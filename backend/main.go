package main

import (
	"github.com/Rocksus/devcamp-2021-big-project/backend/database"
	"github.com/Rocksus/devcamp-2021-big-project/backend/server"
	"github.com/Rocksus/devcamp-2021-big-project/backend/server/handlers/product"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	dbConfig := database.Config{
		User:     "postgres",
		Password: "admin",
		DBName:   "devcamp",
		Port:     5432,
		Host:     "db",
		SSLMode:  "disable",
	}
	db := database.GetDatabaseConnection(dbConfig)

	ph := product.NewProductHandler(db)

	router := mux.NewRouter()

	//registering handlers
	router.HandleFunc("/product", ph.AddProduct).Methods(http.MethodPost)
	router.HandleFunc("/product/{id:[0-9]+}", ph.EditProduct).Methods(http.MethodPut)
	router.HandleFunc("/product/{id:[0-9]+}", ph.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products", ph.GetProductBatch).Methods(http.MethodGet)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9000,
	}
	server.Serve(serverConfig, router)
}
