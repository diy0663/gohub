package cmd

import (
	"time"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/rabbitmq"
	"github.com/spf13/cobra"
)

// todo 最后要保存一次才能自动import ,import 的时候也检查一下是否包引入是正确的

// todo 这个生成的命令 ,还得记得挂到上层命令那里面去

var CmdDelaySend = &cobra.Command{
	Use:   "delay_send",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runDelaySend,
	Args:  cobra.ExactArgs(0), // 只允许且必须传 1 个参数
}

func runDelaySend(cmd *cobra.Command, args []string) {

	console.Success("runDelaySend...")
	mq := rabbitmq.NewRabbitMQDelay("test_delay_exchange")

	defer mq.Destory()

	mq.PublishDelay("this is a delay message", time.Second*time.Duration(15))
	console.Success("runRoutingSend...end")
}
