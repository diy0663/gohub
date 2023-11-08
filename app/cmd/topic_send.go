package cmd

import (
	"strconv"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/rabbitmq"
	"github.com/spf13/cobra"
)

// todo 最后要保存一次才能自动import ,import 的时候也检查一下是否包引入是正确的

// todo 这个生成的命令 ,还得记得挂到上层命令那里面去

var CmdTopicSend = &cobra.Command{
	Use:   "topic_send",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runTopicSend,
	Args:  cobra.ExactArgs(0), // 只允许且必须传 1 个参数
}

func runTopicSend(cmd *cobra.Command, args []string) {

	console.Success("runTopicSend...")
	mq1 := rabbitmq.NewRabbitMQTopic("the_topic_exchange", "key1.one")
	mq2 := rabbitmq.NewRabbitMQTopic("the_topic_exchange", "key2.two")
	defer mq1.Destory()
	defer mq2.Destory()
	for i := 0; i < 10; i++ {
		mq1.PublishTopic("\r\n routing key:" + mq1.Key + "  message  " + strconv.Itoa(i))
		mq2.PublishTopic("\r\n routing key:" + mq2.Key + "  message  " + strconv.Itoa(i))
	}
	console.Success("runTopicSend...end")
}
