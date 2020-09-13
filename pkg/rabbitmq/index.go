package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)


var RabbitMq = struct {
	conn *amqp.Connection
	channel *amqp.Channel
	queueName string
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

// queue 宣告
func QueueDeclare(name string) {
	queue, err := RabbitMq.channel.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		panic("create queue failed")
	}
	RabbitMq.queueName = queue.Name
}

// publish 訊息
func Publish(message string) {
	err := RabbitMq.channel.Publish("", RabbitMq.queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte(message),
	})
	if err != nil {
		fmt.Println("publish error")
	}
}

func Consume(name string) {
	msg, err := RabbitMq.channel.Consume(name, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println("consume message error: ", err)
	} else {
		for d := range msg {
			// TODO: New Lottery
			fmt.Printf("receive message: %s", d.Body)
		}
	}
}