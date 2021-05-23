# Tokopedia Devcamp 2021 Big Project Backend

The backend side will explore various topics ranging from:

- Introduction to Golang
- Server & Database in Golang
- GraphQL
- Monitoring
- Caching
- Message Queueing

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