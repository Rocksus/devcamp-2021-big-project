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
    "name": "product example",
    "description": "An amazing example product!",
    "price": 17500,
    "rating": 0,
    "image_url": "add_this_later",
    "preview_image_url": "add_this_later",
    "slug":"product-example"
}'
```

Getting Product Data

```shell
curl --location --request GET 'http://localhost:9000/product/1'
```

Updating Product

```shell
curl --location --request PUT 'http://localhost:9000/product/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "rating": 5
}'
```

## Code Organization

Code organization will be presented below. Please refer to each folder for a more detailed description.

```
backend
 ├── database                   # Database initialization code
 │     └── database.go
 ├── productmodule              # Handles all of the product data that is used in this project
 │     ├── product.go
 │     ├── query.go
 │     ├── storage.go           # Persistent storage implementation of product
 │     └── type.go
 ├── server                     # REST HTTP server and handlers
 │     ├── handlers
 │     │     └── product        # Product HTTP Handler
 │     ├── server.go
 │     └── template.go
 ├── main.go                    # Main file
 └── Dockerfile                 # Dockerfile that builds the image
```