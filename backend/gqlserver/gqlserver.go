package gqlserver

import (
	"context"
	"fmt"
	"github.com/Rocksus/devcamp-2021-big-project/backend/gqlserver/gql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	Port         int
}

func Serve(cfg Config, schemaWrapper *gql.SchemaWrapper) {

	handler := gql.NewHandler(schemaWrapper)
	http.Handle("/graphql", handler.Handle())

	srv := &http.Server{
		Addr:         fmt.Sprint(":", cfg.Port),
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("[GQLServer] unable to listen and serve, err: " + err.Error())
		}
	}()
	log.Println("[GQLServer] HTTP server is running at port ", cfg.Port)

	s := make(chan os.Signal, 1)

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-s

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println("[GQLServer] error on shutting down HTTP Server, err: ", err.Error())
	}
}
