package notifications

import (
	"../provider"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

var (
	Channel chan notification
)

type notification struct {
	ID         int64                  `json:"id"`
	UserID     int64                  `json:"user_id"`
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
}

func init() {
	Channel = make(chan notification)

	ch, err := provider.AMQPConnection().Channel()
	if err != nil {
		log.Fatal("can't create rabbitmq channel (", err, ")")
	}

	q, err := ch.QueueDeclare("notifications", false, false, false, false, amqp.Table{
		"x-message-ttl": int64(5000),
	})
	if err != nil {
		log.Fatal("can't connect to notification queue (", err, ")")
	}

	go func() {
		for {
			dch, err := ch.Consume(q.Name, "", false, false, false, false, nil)
			if err != nil {
				log.Fatal(err)
			}

			for d := range dch {
				n := notification{}
				err := json.Unmarshal(d.Body, &n)
				if err != nil {
					log.Fatal(err)
				}

				Channel <- n
				d.Ack(false)
			}
		}
	}()
}
