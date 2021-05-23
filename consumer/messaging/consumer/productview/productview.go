package productview

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"log"
)

type Consumer struct {
}

type ProductResponse struct {
	ID              int64   `json:"id,omitempty"`
	Name            string  `json:"name,omitempty"`
	Description     string  `json:"description,omitempty"`
	Price           int64   `json:"price,omitempty"`
	Rating          float32 `json:"rating,omitempty"`
	ImageURL        string  `json:"image_url,omitempty"`
	PreviewImageURL string  `json:"preview_image_url,omitempty"`
	Slug            string  `json:"slug,omitempty"`
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
