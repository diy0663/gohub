package cmd

import (
	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/rabbitmq"
	"github.com/spf13/cobra"
)

// todo 最后要保存一次才能自动import ,import 的时候也检查一下是否包引入是正确的

// todo 这个生成的命令 ,还得记得挂到上层命令那里面去

var CmdTopicReceive = &cobra.Command{
	Use:   "topic_receive",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runTopicReceive,
	Args:  cobra.ExactArgs(0), // 只允许且必须传 1 个参数
}

func runTopicReceive(cmd *cobra.Command, args []string) {

	console.Success("runTopicReceive...")

	//在RabbitMQ的topic模式中，通配符用于匹配消息的路由键（Routing Key）。
	// 通配符有两种类型：前缀通配符“#”和正则通配符“*”。

	// 前缀通配符“#”用于匹配Routing Key中的任意数量的字符（包括0个字符），但必须与前缀完全匹配。
	//例如，如果Routing Key为“foo.bar”，则“#”可以匹配Routing Key为“foo.#”的消

	// 正则通配符“”用于匹配Routing Key中的一个单词或子字符串。
	// 它与前缀通配符不同，可以匹配多个字符，但必须与前缀部分完全匹配。
	// 例如，如果Routing Key为“foo.bar”，则“”可以匹配Routing Key为“foo.*”的消息，但不能匹配Routing Key为“foobar”的消息。

	// args[0] , 可以尝试输入包含通配符的字符串, 看看是否能匹配上 topic_send 那边的 key1.one 跟 key2.two
	// 例如 go run main.go  topic_receive  #
	// 例如 go run main.go  topic_receive  key1.*
	// 例如 go run main.go  topic_receive  #.two

	// #.two 能匹配 key2.two
	//mq := rabbitmq.NewRabbitMQTopic("the_topic_exchange", "#.two")

	// key1.* 能匹配 key1.one
	//mq := rabbitmq.NewRabbitMQTopic("the_topic_exchange", "key1.*")

	// # 能同时匹配 key1.one 和 key2.two
	mq := rabbitmq.NewRabbitMQTopic("the_topic_exchange", "#")

	defer mq.Destory()
	mq.ConsumeTopic()
}
