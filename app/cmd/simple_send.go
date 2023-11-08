package cmd

import (
	"fmt"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/rabbitmq"
	"github.com/spf13/cobra"
)

// todo 最后要保存一次才能自动import ,import 的时候也检查一下是否包引入是正确的

// todo 这个生成的命令 ,还得记得挂到上层命令那里面去

var CmdSimpleSend = &cobra.Command{
	Use:   "simple_send",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runSimpleSend,
	Args:  cobra.ExactArgs(0), // 只允许且必须传 1 个参数
}

func runSimpleSend(cmd *cobra.Command, args []string) {

	console.Success("runSimpleSend....")
	queueName := "SIMPLE_QUEUE_FEFAULT"
	//mq := rabbitmq.NewRabbitMQSimple("SIMPLE_QUEUE_1", rabbitmq.WithUrl("amqp://guest:guest@127.0.0.1:5672/"))
	mq := rabbitmq.NewRabbitMQSimple(queueName)
	defer mq.Destory()
	for i := 0; i <= 10; i++ {
		mq.PublishSimple(fmt.Sprintf("Message No.%v  from  %v \r\n", i, queueName))
	}

}
