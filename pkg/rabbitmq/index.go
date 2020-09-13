package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)


var RabbitMq = struct {
	conn *amqp.Connection
	channel *amqp.Channel
}{}

func init() {
	var err error
	dial := fmt.Sprintf("amqp://%s:%s@%s:%s/",
						os.Getenv("RABBITMQ_USER"),
						os.Getenv("RABBITMQ_PASS"),
						os.Getenv("RABBITMQ_HOST"),
						os.Getenv("RABBITMQ_PORT"))
	RabbitMq.conn, err = amqp.Dial(dial)
	if err != nil {
		fmt.Println("連線失敗")
	}

	RabbitMq.channel, err = RabbitMq.conn.Channel()
	if err != nil {
		fmt.Println("創建 channel 失敗")
	}
}
