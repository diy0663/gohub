package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/rabbitmq"
	"github.com/spf13/cobra"
)

// todo 最后要保存一次才能自动import ,import 的时候也检查一下是否包引入是正确的

// todo 这个生成的命令 ,还得记得挂到上层命令那里面去

var CmdSimpleReceiveHandler = &cobra.Command{
	Use:   "simple_receive_handler",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runSimpleReceiveHandler,
	Args:  cobra.ExactArgs(0), // 只允许且必须传 1 个参数
}

func runSimpleReceiveHandler(cmd *cobra.Command, args []string) {

	console.Success("runSimpleReceive....")
	queueName := "SIMPLE_QUEUE_FEFAULT"
	mq := rabbitmq.NewRabbitMQSimple(queueName)
	defer mq.Destory()
	mq.ConsumeSimpleWithHandler(dealSimpleReceive)
}

// 对消息进行处理 (复杂的参数就是把这个msg 转为结构体)
func dealSimpleReceive(msg string) {
	var data rabbitmq.MsgData
	err := json.Unmarshal([]byte(msg), &data)
	if err != nil {
		console.Error("json.Unmarshal err: " + err.Error())
	} else {
		fmt.Println(" receive: ", data.Time, " by ", data.Msg)
	}
}
