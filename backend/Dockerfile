FROM golang:1.16 AS build

# Environment variables will be put here
ENV GO111MODULE=on
# end

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git

WORKDIR /go/src/devcamp-2021/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .

FROM alpine
COPY --from=build /go/src/devcamp-2021/gqlserver/static go/bin/devcamp-2021/static
COPY --from=build /go/src/devcamp-2021/main go/bin/devcamp-2021/

EXPOSE 9000
WORKDIR go/bin/devcamp-2021/

CMD ["./main"]