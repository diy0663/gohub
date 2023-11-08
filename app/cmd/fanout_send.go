package cmd

import (
	"strconv"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/rabbitmq"
	"github.com/spf13/cobra"
)

// todo 最后要保存一次才能自动import ,import 的时候也检查一下是否包引入是正确的

// todo 这个生成的命令 ,还得记得挂到上层命令那里面去

var CmdFanoutSend = &cobra.Command{
	Use:   "fanout_send",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runFanoutSend,
	Args:  cobra.ExactArgs(0), // 只允许且必须传 1 个参数
}

func runFanoutSend(cmd *cobra.Command, args []string) {

	console.Success("这是一条成功的提示")
	mq := rabbitmq.NewRabbitMQFanout("THE_FANOUT_EXCHANE")
	defer mq.Destory()
	for i := 0; i <= 10; i++ {
		mq.PublishFanout("fanout 广播模式生产 的 第" + strconv.Itoa(i) + "条" + "数据 \r\n")
	}

}
