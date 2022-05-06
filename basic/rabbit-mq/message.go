package rabbit_mq

import (
	"github.com/streadway/amqp"
)

type MessageClient struct {
	conn *amqp.Connection
}

type IMessageClient interface {
	NewConnection(connectionStr string) error
	PublishToQueue(data []byte, queueName string) error
	SubscribeToQueue(queueName string, handlerFunc func(delivery amqp.Delivery)) error
	Close()
}

func (m *MessageClient) NewConnection(connectionStr string) error {
	if connectionStr == "" {
		panic("the connection str mustnt be null")
	}
	var err error
	m.conn, err = amqp.Dial(connectionStr)
	return err
}

func (m *MessageClient) PublishToQueue(data []byte, queueName string) error {
	if m.conn == nil {
		panic("before publish u must connect the RabbitMq first")
	}
	ch, err := m.conn.Channel()
	defer ch.Close()
	if err != nil {
		panic(err)
	}

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		panic(err)
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		})
	if err != nil {
		return err
	}
	return nil
}

func (m *MessageClient) SubscribeToQueue(queueName string, handlerFunc func(delivery amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	//defer ch.Close()
	if err != nil {
		panic(err)
	}

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		panic(err)
	}

	go consumeLoop(msgs, handlerFunc)
	return nil
}

func (m *MessageClient) Close() {
	defer m.Close()
}

func consumeLoop(msg <-chan amqp.Delivery, handlerFunc func(delivery amqp.Delivery)) {
	for d := range msg {
		handlerFunc(d)
	}
}
