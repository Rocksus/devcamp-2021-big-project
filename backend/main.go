package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Rocksus/devcamp-2021-big-project/backend/cache"

	"github.com/gorilla/mux"

	"github.com/Rocksus/devcamp-2021-big-project/backend/database"
	"github.com/Rocksus/devcamp-2021-big-project/backend/gqlserver/gql"
	"github.com/Rocksus/devcamp-2021-big-project/backend/gqlserver/gql/product"
	"github.com/Rocksus/devcamp-2021-big-project/backend/productmodule"
	"github.com/Rocksus/devcamp-2021-big-project/backend/server"
	productHandler "github.com/Rocksus/devcamp-2021-big-project/backend/server/handlers/product"
)

func main() {
	cacheConfig := cache.Config{
		MaxActive:       200,
		MaxIdle:         5,
		IdleTimeout:     3 * time.Second,
		MaxConnLifetime: 0,
		Address:         "redis:6379",
	}

	cache := cache.InitializeRedis(cacheConfig)

	dbConfig := database.Config{
		User:     "postgres",
		Password: "admin",
		DBName:   "devcamp",
		Port:     5432,
		Host:     "db",
		SSLMode:  "disable",
	}
	db := database.GetDatabaseConnection(dbConfig)

	pm := productmodule.NewProductModule(db, cache)
	ph := productHandler.NewProductHandler(pm)
	pr := product.NewResolver(pm)

	schemaWrapper := gql.NewSchemaWrapper().
		WithProductResolver(pr)

	if err := schemaWrapper.Init(); err != nil {
		log.Fatal("unable to parse schema, err: ", err.Error())
	}

	router := mux.NewRouter()

	// REST Handlers
	router.HandleFunc("/product", ph.AddProduct).Methods(http.MethodPost)
	router.HandleFunc("/product/{id:[0-9]+}", ph.UpdateProduct).Methods(http.MethodPut)
	router.HandleFunc("/product/{id:[0-9]+}", ph.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products", ph.GetProductBatch).Methods(http.MethodGet)

	// GraphQL Handler
	router.Path("/graphql").Handler(gql.NewHandler(schemaWrapper).Handle())

	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/").Handler(fs)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9000,
	}
	server.Serve(serverConfig, router)
}
