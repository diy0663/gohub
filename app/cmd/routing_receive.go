package cmd

import (
	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/rabbitmq"
	"github.com/spf13/cobra"
)

// todo 最后要保存一次才能自动import ,import 的时候也检查一下是否包引入是正确的

// todo 这个生成的命令 ,还得记得挂到上层命令那里面去

var CmdRoutingReceive = &cobra.Command{
	Use:   "routing_receive",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runRoutingReceive,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runRoutingReceive(cmd *cobra.Command, args []string) {

	console.Success("runRoutingReceive....")
	mq := rabbitmq.NewRabbitMQRouting("the_routing_exchange", args[0])
	defer mq.Destory()
	mq.ConsumeRouting()
}
