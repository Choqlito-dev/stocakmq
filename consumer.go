package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMQ) Consumir(queue string, handler func(amqp.Delivery)) error {
	msgs, err := r.channel.Consume(
		queue,
		"",
		true,  // autoAck
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			handler(msg)
		}
	}()

	return nil
}
