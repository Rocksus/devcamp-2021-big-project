package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type Config struct {
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	Port         int
}

func Serve(cfg Config, router *mux.Router) {
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprint(":", cfg.Port),
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("[Server] unable to listen and serve, err: " + err.Error())
		}
	}()
	log.Println("[Server] HTTP server is running at port ", cfg.Port)

	s := make(chan os.Signal, 1)

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-s

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println("[Server] error on shutting down HTTP Server, err: ", err.Error())
	}
}
