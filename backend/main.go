package main

import (
	"github.com/Rocksus/devcamp-2021-big-project/backend/database"
	"github.com/Rocksus/devcamp-2021-big-project/backend/gqlserver"
	"github.com/Rocksus/devcamp-2021-big-project/backend/gqlserver/gql"
	"github.com/Rocksus/devcamp-2021-big-project/backend/gqlserver/gql/product"
	"log"
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

	ps := product.NewProductService(db)
	productResolver := product.NewResolver(ps)

	schemaWrapper := gql.NewSchemaWrapper().
		WithProductResolver(productResolver)

	if err := schemaWrapper.Init(); err != nil {
		log.Fatal("unable to parse schema, err: ", err.Error())
	}

	serverConfig := gqlserver.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9000,
	}
	gqlserver.Serve(serverConfig, schemaWrapper)
}
