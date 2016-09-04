package provider

import (
	"github.com/streadway/amqp"
	"gopkg.in/redis.v3"
	"log"
	"sync"
)

var (
	amqpConnection *amqp.Connection
	redisClient    *redis.Client
)

func init() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		var err error
		amqpConnection, err = amqp.Dial("amqp://rabbitmq:rabbitmq@localhost:5672/")
		if err != nil {
			log.Fatal("can't connect to rabbitmq (", err, ")")
		}

		wg.Done()
	}()

	go func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		_, err := redisClient.Ping().Result()
		if err != nil {
			log.Fatal("can't connect to redis (", err, ")")
		}

		wg.Done()
	}()

	wg.Wait()
}

func AMQPConnection() *amqp.Connection {
	return amqpConnection
}

func RedisClient() *redis.Client {
	return redisClient
}
