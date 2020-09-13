package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"go_rabbitmq_bet/pkg/setting"
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

	// 註冊 queue
	for _, name := range setting.QueueName {
		QueueDeclare(name)
	}
}

// queue 宣告
func QueueDeclare(name string) {
	fmt.Println("queue-" + name, " register ....")
	_, err := RabbitMq.channel.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		panic("create queue failed")
	}
}

/**
	publish 訊息
	@string queueName - 隊列名稱
	@string message - 訊息
 */
func Publish(queueName string, message interface{}) {
	b, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("message json encode error")
	}
	err = RabbitMq.channel.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body: b,
	})
	if err != nil {
		fmt.Println("publish error")
	}
}

/**
	取得隊列 訊息
	@param string queueName - 隊列名稱
 */
func GetQueueMessage(queueName string) *amqp.Delivery{
	msg, ok, err := RabbitMq.channel.Get(queueName, true)
	if err != nil {
		fmt.Println("consume message error: ", err)
		return nil
	}

	if !ok {
		fmt.Println("do not get message")
		return nil
	}
	return &msg
}

/**
	消費queue
	@string queueName - 隊列名稱
 */
func ConsumeBets(queueName string){
	msg := GetQueueMessage(queueName)
	if msg != nil {
		fmt.Println("consume bet")
	}
}