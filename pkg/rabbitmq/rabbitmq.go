package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// todo 应该改为.env 配置化
const MQURL = "amqp://guest:guest@127.0.0.1:5672/vhost1"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	// binding key ,用于路由匹配
	Key string
	// 连接配置信息
	Mqurl string
}

type RabbitMQOption func(*RabbitMQ)

func WithUrl(url string) RabbitMQOption {
	return func(r *RabbitMQ) {
		r.Mqurl = url
	}
}

// 其实仅仅是传参而已,并没有正常去生成一个rabbitmq连接实例
func newRabbitMQ(queueName, exchangeName, key string, options ...RabbitMQOption) *RabbitMQ {
	r := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchangeName,
		Key:       key,
	}
	for _, option := range options {
		option(r)
	}
	// 没传的默认值
	if r.Mqurl == "" {
		r.Mqurl = MQURL
	}
	return r

}

func (r *RabbitMQ) Destory() error {
	err := r.channel.Close()
	if err != nil {
		return err
	}
	err = r.conn.Close()
	if err != nil {
		return err
	}
	return nil

}

func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err.Error())
		panic(fmt.Sprintf("%s:%s", message, err.Error()))
	}
}

// 配置好conn 以及channel之后 ,里面没有声明queue
func NewRabbitMQSimple(queueName string, options ...RabbitMQOption) *RabbitMQ {
	// 简单模式, 不需要交换机以及route key
	r := newRabbitMQ(queueName, "", "", options...)

	var err error
	r.conn, err = amqp.Dial(r.Mqurl)
	r.failOnErr(err, "Failed to connect to RabbitMQ")
	r.channel, err = r.conn.Channel()
	r.failOnErr(err, "Failed to open a channel")
	return r

}

// queue 的声明跟使用(发送或者接收消息) 都在一起做处理

func (r *RabbitMQ) PublishSimple(message string) {

	// 声明创建队列

	_, err := r.channel.QueueDeclare(
		r.QueueName, // name
		false,       // durable 持久化
		false,       // delete when unused 不用时删除
		false,       // exclusive 是否独占
		false,       // no-wait 是否立即
		nil,         // arguments
	)
	r.failOnErr(err, "Failed to declare a queue when PublishSimple ")
	err = r.channel.Publish(
		"",          // exchange
		r.QueueName, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	r.failOnErr(err, "Failed to publish a message when PublishSimple")
}

func (r *RabbitMQ) ConsumeSimple() {
	q, err := r.channel.QueueDeclare(
		r.QueueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	r.failOnErr(err, "Failed to declare a queue when ConsumeSimple")
	messages, err := r.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack  自动应答
		false,  // exclusive
		false,  // no-local  true则表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false,  // no-wait
		nil,    // args
	)
	r.failOnErr(err, "Failed to register a consumer when ConsumeSimple")
	forever := make(chan bool)
	go func() {
		for d := range messages {
			fmt.Printf("Received a message: %s", d.Body)
		}
	}()
	// 阻塞在这里
	<-forever
}
