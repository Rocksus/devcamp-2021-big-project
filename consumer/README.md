# Tokopedia Devcamp 2021 Big Project Consumer

The consumer side will be used to receive messages from the backend producer.

## Code Organization

Code organization will be presented below. Please refer to each folder for a more detailed description.

```
consumer
 ├── messaging          # Messaging Producer & Consumer
 │     ├── consumer.go
 │     └── producer.go
 ├── tracer             # Opentracing & Jaeger interactor
 │     └── tracer.go
 ├── main.go            # Main file
 └── Dockerfile         # Dockerfile that builds the image
```