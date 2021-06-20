package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/Rocksus/devcamp-2021-big-project/backend/database"
	"github.com/Rocksus/devcamp-2021-big-project/backend/productmodule"
	"github.com/Rocksus/devcamp-2021-big-project/backend/server"
	productHandler "github.com/Rocksus/devcamp-2021-big-project/backend/server/handlers/product"
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

	pm := productmodule.NewProductModule(db)
	ph := productHandler.NewProductHandler(pm)

	router := mux.NewRouter()

	// REST Handlers
	router.HandleFunc("/product", ph.AddProduct).Methods(http.MethodPost)
	router.HandleFunc("/product/{id:[0-9]+}", ph.UpdateProduct).Methods(http.MethodPut)
	router.HandleFunc("/product/{id:[0-9]+}", ph.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products", ph.GetProductBatch).Methods(http.MethodGet)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9000,
	}
	server.Serve(serverConfig, router)
}
