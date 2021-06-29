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

### REST

------

Adding New Product

```shell
curl --location --request POST 'http://localhost:9000/product' \
--header 'Content-Type: application/json' \
--data-raw '{
    "product_name": "product example",
    "product_description": "An amazing example product!",
    "product_price": 17500,
    "rating": 0,
    "product_image": "https://images.tokopedia.net/img/cache/900/product-1/2020/7/2/16620763/16620763_b0f98181-2092-4a28-9035-d588efd495c1_1000_1000",
    "additional_product_image": [
        "https://images.tokopedia.net/img/cache/900/product-1/2020/7/2/16620763/16620763_b75ad660-02f8-475c-bd3e-278d9205fbc2_1000_1000"
    ]
}'
```

Getting Product Data

```shell
curl --location --request GET 'http://localhost:9000/product/1'
```

Getting Multiple Products

```shell
curl --location --request GET 'http://localhost:9000/products?limit=10&offset=0'
```

Updating Product

```shell
curl --location --request PUT 'http://localhost:9000/product/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "rating": 5
}'
```

### GraphQL

------

Adding New Product

```shell
curl --location --request POST 'http://localhost:9000/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{
    "query": "mutation { AddProduct(product_name:\"cool new product\" product_description:\"This is a cool new product. Must have if you are trendy\" product_price:299000 rating:5 product_image:\"https://images.tokopedia.net/img/cache/900/product-1/2020/7/2/16620763/16620763_b0f98181-2092-4a28-9035-d588efd495c1_1000_1000\") { product_id } }"
}'
```

Getting Product Data

```shell
curl --location --request POST 'http://localhost:9000/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{
    "query": "query { ProductDetail(productId:1) {product_id product_name product_description product_price rating product_image additional_product_image} }"
}'
```

Getting Multiple Products

```shell
curl --location --request POST 'http://localhost:9000/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{
    "query": "query { ProductLists(limit:2, offset:1) {product_id product_name product_description product_price rating product_image additional_product_image} }"
}'
```

Updating Product

```shell
curl --location --request POST 'http://localhost:9000/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{
    "query": "mutation { UpdateProduct(product_id:1 product_price:499000) { product_id } }"
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
 │     └── static               # GraphQL Playground assets
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