package rabbitmq

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// todo 应该改为.env 配置化
const MQURL = "amqp://guest:guest@127.0.0.1:5672/vhost1"

type HandlerFunc func(message string)

// 定义一个基本的队列传输消息结构体
type MsgData struct {
	Msg  string `json:"msg,omitempty" `
	Time string `json:"time,omitempty" `
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	// binding key ,用于路由匹配
	Key string
	// 连接配置信息
	Mqurl            string
	IsDurable        bool
	DeliveryMode     uint8
	IsConsumeAutoAck bool
}

type RabbitMQOption func(*RabbitMQ)

func WithUrl(url string) RabbitMQOption {
	return func(r *RabbitMQ) {
		r.Mqurl = url
	}
}

func WithComsumeAutoAck(autoAck bool) RabbitMQOption {
	return func(r *RabbitMQ) {
		r.IsConsumeAutoAck = autoAck
	}
}

func WithDurable(durable bool) RabbitMQOption {
	return func(r *RabbitMQ) {
		r.IsDurable = durable
	}
}

// 其实仅仅是传参而已,并没有正常去生成一个rabbitmq连接实例
func newRabbitMQ(queueName, exchangeName, key string, options ...RabbitMQOption) *RabbitMQ {
	r := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchangeName,
		Key:       key,
		// 默认持久化
		IsDurable:    true,
		DeliveryMode: 2,
		// 默认自动应答
		IsConsumeAutoAck: true,
	}
	for _, option := range options {
		option(r)
	}
	// 没传的默认值
	if r.Mqurl == "" {
		r.Mqurl = MQURL
	}
	// 持久化判断
	if r.IsDurable {
		r.DeliveryMode = 2
	} else {
		r.DeliveryMode = 1
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
		r.IsDurable, // durable 持久化
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
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: r.DeliveryMode,
		})
	r.failOnErr(err, "Failed to publish a message when PublishSimple")
}

func (r *RabbitMQ) ConsumeSimple() {
	q, err := r.channel.QueueDeclare(
		r.QueueName, // name
		r.IsDurable, // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	r.failOnErr(err, "Failed to declare a queue when ConsumeSimple")

	// 在一条消息没被确认处理完之前,不消费新的消息
	// 设置用于控制消费者从队列中获取消息的速率,均衡worker的工作量
	err = r.channel.Qos(
		1, 0, false,
	)
	r.failOnErr(err, "Failed to set Qos")

	messages, err := r.channel.Consume(
		q.Name,             // queue
		"",                 // consumer
		r.IsConsumeAutoAck, // auto-ack  自动应答
		false,              // exclusive
		false,              // no-local  true则表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false,              // no-wait
		nil,                // args
	)
	r.failOnErr(err, "Failed to register a consumer when ConsumeSimple")
	forever := make(chan bool)
	go func() {
		for d := range messages {
			fmt.Printf("Received a message: %s", d.Body)
			if !r.IsConsumeAutoAck {
				d.Ack(false)
			}
		}
	}()
	// 阻塞在这里
	<-forever
}

func (r *RabbitMQ) ConsumeSimpleWithHandler(hanler HandlerFunc) {
	q, err := r.channel.QueueDeclare(
		r.QueueName, // name
		r.IsDurable, // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	r.failOnErr(err, "Failed to declare a queue when ConsumeSimple")

	// 在一条消息没被确认处理完之前,不消费新的消息
	// 设置用于控制消费者从队列中获取消息的速率,均衡worker的工作量
	err = r.channel.Qos(
		1, 0, false,
	)
	r.failOnErr(err, "Failed to set Qos")

	messages, err := r.channel.Consume(
		q.Name,             // queue
		"",                 // consumer
		r.IsConsumeAutoAck, // auto-ack  自动应答
		false,              // exclusive
		false,              // no-local  true则表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false,              // no-wait
		nil,                // args
	)
	r.failOnErr(err, "Failed to register a consumer when ConsumeSimple")
	forever := make(chan bool)
	go func() {
		for d := range messages {
			s := BytesToString(&d.Body)
			hanler(*s)
			if !r.IsConsumeAutoAck {
				d.Ack(false)
			}
		}
	}()
	// 阻塞在这里
	<-forever
}

// 广播模式
func NewRabbitMQFanout(exchangeName string, options ...RabbitMQOption) *RabbitMQ {
	r := newRabbitMQ("", exchangeName, "", options...)
	var err error
	r.conn, err = amqp.Dial(r.Mqurl)
	r.failOnErr(err, "Failed to connect to RabbitMQ")
	r.channel, err = r.conn.Channel()
	r.failOnErr(err, "Failed to open a channel")
	return r
}

// 广播模式的生产者
func (r *RabbitMQ) PublishFanout(message string) {
	// 声明交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,  // name
		"fanout",    // type
		r.IsDurable, // durable 持久化
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	r.failOnErr(err, "Failed to declare an exchange when PublishFanout")
	// 发送广播消息
	err = r.channel.Publish(
		r.Exchange, // exchange
		"",         // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: r.DeliveryMode,
		})
	r.failOnErr(err, "Failed to publish a message when PublishFanout")
}

// 广播模式的订阅接收者
func (r *RabbitMQ) ConsumeFanout() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,  // name
		"fanout",    // type
		r.IsDurable, // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange when ConsumeFanout:%s", err.Error())
	}

	// 创建队列,注意这里的队列不需要写名称
	q, err := r.channel.QueueDeclare(
		"",          // name ,随机生产队列名称
		r.IsDurable, // durable
		false,       // delete when unused
		true,        // exclusive 独占
		false,       // no-wait
		nil,         // arguments
	)
	r.failOnErr(err, "Failed to declare a queue when ConsumeFanout")
	// 绑定队列到 交换机
	err = r.channel.QueueBind(
		q.Name, // queue name
		"",     // routing key  广播模式不需要route key
		r.Exchange,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind a queue with exchange when ConsumeFanout:%s", err.Error())
	}

	// 消费消息
	messages, err := r.channel.Consume(
		q.Name,             // queue
		"",                 // consumer
		r.IsConsumeAutoAck, // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer when ConsumeFanout:%s", err.Error())
	}
	forever := make(chan bool)
	go func() {
		for message := range messages {
			fmt.Printf("Received a message: %s", message.Body)
			if !r.IsConsumeAutoAck {
				message.Ack(false)
			}
		}
	}()

	// 阻塞在这里
	<-forever
}

// 路由模式,在交换器上做完全匹配之后转发消息到符合条件的队列上去
func NewRabbitMQRouting(exchangeName, key string, options ...RabbitMQOption) *RabbitMQ {
	r := newRabbitMQ("", exchangeName, key, options...)
	var err error
	r.conn, err = amqp.Dial(r.Mqurl)
	r.failOnErr(err, "Failed to connect to RabbitMQ")
	r.channel, err = r.conn.Channel()
	r.failOnErr(err, "Failed to open a channel")
	return r
}

func (r *RabbitMQ) PublishRouting(message string) {
	// 1. 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,  // name
		"direct",    // type 注意 ,路由模式就得是 direct 类型
		r.IsDurable, // durable 持久化
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	r.failOnErr(err, "Failed to declare an exchange when PublishRouting")

	// 2.发送消息
	err = r.channel.Publish(
		r.Exchange, // exchange
		r.Key,      // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: r.DeliveryMode,
		})
	r.failOnErr(err, "Failed to publish a message when PublishRouting")
}

func (r *RabbitMQ) ConsumeRouting() {
	// 创建交换器
	err := r.channel.ExchangeDeclare(
		r.Exchange,  // name
		"direct",    // type
		r.IsDurable, // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	r.failOnErr(err, "Failed to declare an exchange when ConsumeRouting")

	// 创建队列
	q, err := r.channel.QueueDeclare(
		"",          // name,随机生产队列名称
		r.IsDurable, // durable
		false,       // delete when unused
		true,        // exclusive 独占
		false,       // no-wait
		nil,         // arguments
	)
	r.failOnErr(err, "Failed to declare a queue when ConsumeRouting")

	// 绑定队列到exchange
	err = r.channel.QueueBind(
		q.Name, // queue name
		r.Key,  // routing key
		r.Exchange,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to bind a queue with exchange when ConsumeRouting")
	messages, err := r.channel.Consume(
		q.Name,             // queue
		"",                 // consumer
		r.IsConsumeAutoAck, // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	r.failOnErr(err, "Failed to register a consumer when ConsumeRouting")
	forever := make(chan bool)
	go func() {
		for message := range messages {
			fmt.Printf(" the key "+r.Key+"  Received a message: %s\r\n", message.Body)
			if !r.IsConsumeAutoAck {
				message.Ack(false)
			}
		}
	}()
	<-forever

}

// topic 模式,支持模糊匹配
func NewRabbitMQTopic(exchangeName, key string, options ...RabbitMQOption) *RabbitMQ {
	r := newRabbitMQ("", exchangeName, key, options...)
	var err error
	r.conn, err = amqp.Dial(r.Mqurl)
	r.failOnErr(err, "Failed to connect to RabbitMQ")
	r.channel, err = r.conn.Channel()
	r.failOnErr(err, "Failed to open a channel")
	return r
}

func (r *RabbitMQ) PublishTopic(message string) {
	// 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,  // name
		"topic",     // type
		r.IsDurable, // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	r.failOnErr(err, "Failed to declare an exchange when PublishTopic")

	err = r.channel.Publish(
		r.Exchange, // exchange
		r.Key,      // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: r.DeliveryMode,
		})
	r.failOnErr(err, "Failed to publish a message when PublishTopic")

}

func (r *RabbitMQ) ConsumeTopic() {
	// 创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,  // name
		"topic",     // type
		r.IsDurable, // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	r.failOnErr(err, "Failed to declare an exchange when ConsumeTopic")
	// 创建消息队列
	q, err := r.channel.QueueDeclare(
		"",          // name,随机生产队列名称
		r.IsDurable, // durable
		false,       // delete when unused
		true,        // exclusive 独占
		false,       // no-wait
		nil,         // arguments
	)
	r.failOnErr(err, "Failed to declare a queue when ConsumeTopic")

	// 绑定队列到交换机
	r.channel.QueueBind(
		q.Name, // queue name
		r.Key,  // routing key
		r.Exchange,
		false,
		nil)

	// 消费消息
	messages, err := r.channel.Consume(
		q.Name,             // queue
		"",                 // consumer
		r.IsConsumeAutoAck, // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	r.failOnErr(err, "Failed to register a consumer when ConsumeTopic")
	forever := make(chan bool)
	go func() {
		for message := range messages {
			fmt.Printf("\r\nReceived a message: %s\r\n", message.Body)
			if !r.IsConsumeAutoAck {
				message.Ack(false)
			}
		}
	}()
	<-forever
}

func NewRabbitMQDelay(exchangeName string, options ...RabbitMQOption) *RabbitMQ {
	r := newRabbitMQ("", exchangeName, "", options...)
	var err error
	r.conn, err = amqp.Dial(r.Mqurl)
	r.failOnErr(err, "Failed to connect to RabbitMQ")
	r.channel, err = r.conn.Channel()
	r.failOnErr(err, "Failed to open a channel")
	return r
}

func (r *RabbitMQ) PublishDelay(message string, delay time.Duration) {
	// 1. 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,          // name
		"x-delayed-message", // type todo 重点
		true,                // durable 持久化
		true,                // auto-deleted
		false,               // internal
		false,               // no-wait
		amqp.Table{
			"x-delayed-type": "direct", // todo 重点
		}, // arguments
	)
	r.failOnErr(err, "Failed to declare an exchange when PublishDelay")

	err = r.channel.Publish(
		r.Exchange, // exchange
		"",         // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: r.DeliveryMode,
			Headers: amqp.Table{
				"x-delay": delay.Milliseconds(), // todo 重点
			},
		})
	r.failOnErr(err, "Failed to publish a message when PublishDelay")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " [x-delay] pamas", delay)

}

func (r *RabbitMQ) ConsumeDelay() {
	// 创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,          // name
		"x-delayed-message", // type todo 重点
		true,                // durable
		true,                // auto-deleted
		false,               // internal
		false,               // no-wait
		amqp.Table{
			"x-delayed-type": "direct", // todo 重点
		}, // arguments
	)
	r.failOnErr(err, "Failed to declare an exchange when ConsumeDelay")
	q, err := r.channel.QueueDeclare(
		"",    // name,随机生产队列名称
		true,  // durable
		true,  // delete when unused
		true,  // exclusive 独占
		false, // no-wait
		nil,   // arguments
	)
	r.failOnErr(err, "Failed to declare a queue when ConsumeDelay")

	// 绑定队列到交换机
	err = r.channel.QueueBind(
		q.Name, // queue name
		r.Key,  // routing key
		r.Exchange,
		false,
		nil)
	r.failOnErr(err, "Failed to bind a queue with exchange when ConsumeDelay")

	// 消费消息
	messages, err := r.channel.Consume(
		q.Name,             // queue
		"",                 // consumer
		r.IsConsumeAutoAck, // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	r.failOnErr(err, "Failed to register a consumer when ConsumeDelay")
	forever := make(chan bool)
	go func() {
		for message := range messages {
			fmt.Printf("%v \r\nReceived a message: %s\r\n", time.Now().Format("2006-01-02 15:04:05"), message.Body)
			if !r.IsConsumeAutoAck {
				message.Ack(false)
			}
		}
	}()
	<-forever

}
