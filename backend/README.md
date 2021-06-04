# Tokopedia Devcamp 2021 Big Project Backend

The backend side will explore various topics ranging from:

- Introduction to Golang
- Server & Database in Golang
- GraphQL
- Monitoring
- Caching
- Message Queueing

## Starting Up

If you have ran the `docker-compose` with the `detached` option like so:

```shell
docker-compose up -d
```

you can simply run

```shell
docker-compose up --build app
```

to rebuild the backend service.

## Basic Queries


Do note that we also provide an accessible GraphiQL playground that can be accessed by going to `localhost:9000` (the default port)\
Here we list some example queries for you to start with:

Adding New Product

```shell
curl --location --request POST 'http://localhost:9000/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{
    "query": "mutation { addProduct(name:\"cool new product\" description:\"This is a cool new product. Must have if you are trendy\" price:299000 rating:5 slug:\"cool-new-product\") { id } }"
}'
```

Getting Product Data

```shell
curl --location --request POST 'http://localhost:9000/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{
    "query": "query { product(id:1) { id name description price rating imageURL previewImageURL slug} }"
}'
```

Getting Multiple Products

```shell
curl --location --request POST 'http://localhost:9000/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{
    "query": "query { products(limit:5) { id name description price rating imageURL previewImageURL slug} }"
}'
```

Updating Product

```shell
curl --location --request POST 'http://localhost:9000/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{
    "query": "mutation { updateProduct(id:1 price:499000) { id } }"
}'
```

## Code Organization

Code organization will be presented below. Please refer to each folder for a more detailed description.

```
backend
 ├── cache                      # Cache initialization code
 │     └── cache.go
 ├── database                   # Database initialization code
 │     └── database.go
 ├── gqlserver                  # GQLServer & GQL Resolvers
 │     ├── gql
 │     │     ├── product        # Product resolver
 │     │     ├── schema.go
 │     │     ├── schema.graphql
 │     │     └── handler.go
 │     ├── static               # GraphQL Playground assets
 │     └── gqlserver.go
 ├── messaging                  # Messaging Producer & Consumer
 │     ├── consumer.go
 │     └── producer.go
 ├── monitoring                 # Monitoring interactor
 │     ├── metrics.go           # Defined metrics that are used in this project
 │     ├── middleware.go
 │     └── monitoring.go
 ├── productmodule              # Handles all of the product data that is used in this project
 │     ├── product.go
 │     ├── query.go
 │     ├── cache.go             # Cache implementation of product
 │     ├── storage.go           # Persistent storage implementation of product
 │     └── type.go
 ├── server                     # REST HTTP server and handlers
 │     ├── handlers
 │     │     └── product        # Product HTTP Handler
 │     ├── server.go
 │     └── template.go
 ├── tracer                     # Opentracing & Jaeger interactor
 │     └── tracer.go
 ├── main.go                    # Main file
 └── Dockerfile                 # Dockerfile that builds the image
```