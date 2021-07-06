package productview

import (
	"encoding/json"
	"log"

	"github.com/nsqio/go-nsq"
)

type Consumer struct {
}

type ProductResponse struct {
	ID                 int64    `json:"product_id,omitempty"`
	Name               string   `json:"product_name,omitempty"`
	Description        string   `json:"product_description,omitempty"`
	Price              int64    `json:"product_price,omitempty"`
	PriceFormat        string   `json:"product_price_format,omitempty"`
	Rating             float32  `json:"rating,omitempty"`
	ImageURL           string   `json:"product_image,omitempty"`
	AdditionalImageURL []string `json:"additional_product_image,omitempty"`
}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (pvc *Consumer) HandleMessage(message *nsq.Message) error {
	var msg struct {
		Event         string          `json:"event"`
		ProductDetail ProductResponse `json:"product_detail"`
	}

	if err := json.Unmarshal(message.Body, &msg); err != nil {
		return err
	}

	log.Println("[Consumer] Got event " + msg.Event)
	log.Println("Details: ", msg.ProductDetail)

	return nil
}
