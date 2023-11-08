package cmd

import (
	"strconv"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/rabbitmq"
	"github.com/spf13/cobra"
)

// todo 最后要保存一次才能自动import ,import 的时候也检查一下是否包引入是正确的

// todo 这个生成的命令 ,还得记得挂到上层命令那里面去

var CmdRoutingSend = &cobra.Command{
	Use:   "routing_send",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runRoutingSend,
	Args:  cobra.ExactArgs(0),
}

func runRoutingSend(cmd *cobra.Command, args []string) {

	console.Success("runRoutingSend...")
	mq1 := rabbitmq.NewRabbitMQRouting("the_routing_exchange", "key1")
	mq2 := rabbitmq.NewRabbitMQRouting("the_routing_exchange", "key2")
	defer mq1.Destory()
	defer mq2.Destory()
	for i := 0; i < 10; i++ {
		mq1.PublishRouting("\r\n routing key:" + mq1.Key + "  message  " + strconv.Itoa(i))
		mq2.PublishRouting("\r\n routing key:" + mq2.Key + "  message  " + strconv.Itoa(i))
	}
	console.Success("runRoutingSend...end")

}
