package messaging

import (
	"log"

	"github.com/nsqio/go-nsq"
)

type (
	ConsumerConfig struct {
		Topic         string
		Channel       string
		LookupAddress string
		MaxAttempts   uint16
		MaxInFlight   int
		Handler       nsq.HandlerFunc
	}

	Consumer struct {
		cons          *nsq.Consumer
		lookupAddress string
		handler       nsq.HandlerFunc
	}
)

func NewConsumer(cfg ConsumerConfig) Consumer {
	nsqConf := nsq.NewConfig()
	nsqConf.MaxAttempts = cfg.MaxAttempts
	nsqConf.MaxInFlight = cfg.MaxInFlight

	topic := cfg.Topic
	c, err := nsq.NewConsumer(topic, cfg.Channel, nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}
	return Consumer{
		cons:          c,
		lookupAddress: cfg.LookupAddress,
		handler:       cfg.Handler,
	}
}

func (c *Consumer) Run() {
	c.cons.AddHandler(c.handler)
	err := c.cons.ConnectToNSQLookupd(c.lookupAddress)
	if err != nil {
		log.Fatal(err)
	}
}
