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