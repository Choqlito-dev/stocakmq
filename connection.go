package rabbitmq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// nueva conexion crea una conexión a RabbitMQ
func NuevaConexion(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("error conectando a RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("error creando canal: %w", err)
	}

	return &RabbitMQ{conn: conn, channel: ch}, nil
}

// cerrar conexión y canal
func (r *RabbitMQ) Cerrar() {
	if err := r.channel.Close(); err != nil {
		log.Println("Error cerrando canal:", err)
	}
	if err := r.conn.Close(); err != nil {
		log.Println("Error cerrando conexión:", err)
	}
}
