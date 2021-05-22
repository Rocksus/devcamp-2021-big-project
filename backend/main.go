package main

import (
	"github.com/Rocksus/devcamp-2021-big-project/backend/database"
	"github.com/Rocksus/devcamp-2021-big-project/backend/gqlserver"
	"github.com/Rocksus/devcamp-2021-big-project/backend/gqlserver/gql"
	"github.com/Rocksus/devcamp-2021-big-project/backend/gqlserver/gql/product"
	"github.com/Rocksus/devcamp-2021-big-project/backend/monitoring"
	"github.com/Rocksus/devcamp-2021-big-project/backend/productmodule"
	"github.com/Rocksus/devcamp-2021-big-project/backend/tracer"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"time"
)

func main() {
	if err := monitoring.Init(); err != nil {
		log.Fatal("unable to init monitoring, err: ", err.Error())
	}

	closer, err := tracer.Init("devcamp-backend")
	if err != nil {
		log.Fatal("unable to init tracer, err: ", err.Error())
	}
	defer closer.Close()

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
	productResolver := product.NewResolver(pm)



	schemaWrapper := gql.NewSchemaWrapper().
		WithProductResolver(productResolver)

	if err := schemaWrapper.Init(); err != nil {
		log.Fatal("unable to parse schema, err: ", err.Error())
	}

	router := mux.NewRouter()
	router.Use(monitoring.Middleware)

	router.Path("/graphql").Handler(gql.NewHandler(schemaWrapper).Handle())
	router.Path("/prometheus").Handler(promhttp.Handler())

	serverConfig := gqlserver.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9000,
	}
	gqlserver.Serve(serverConfig, router)
}
