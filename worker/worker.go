package main

import (
	"github.com/Shopify/sarama"
)

func connectConsumer(brokerUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// create new consumer
	conn, err := sarama.NewConsumer(brokerUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}