package main

import (
	"github.com/Rocksus/devcamp-2021-big-project/consumer/messaging"
	"github.com/Rocksus/devcamp-2021-big-project/consumer/messaging/consumer/productview"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	channelProductViewConsumer = "devcamp_productview_consumer"
	topicProductView           = "product_view"
)

func main() {
	pvc := productview.NewConsumer()

	consumerCfg := messaging.ConsumerConfig{
		Channel:       channelProductViewConsumer,
		LookupAddress: "nsqlookupd:4161",
		Topic:         topicProductView,
		MaxAttempts:   10,
		MaxInFlight:   100,
		Handler:       pvc.HandleMessage,
	}

	consumer := messaging.NewConsumer(consumerCfg)

	consumer.Run()

	// keep app alive until terminated
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case <-term:
		log.Println("Application terminated")
	}
}
