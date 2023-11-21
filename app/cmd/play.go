package cmd

import (
	"context"
	"fmt"

	"github.com/diy0663/gohub/app/models/category"
	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runPlay,
	Args:  cobra.ExactArgs(0), // 只允许且必须传 1 个参数
}

func runPlay(cmd *cobra.Command, args []string) {

	console.Warning("开始执行")

	//err := category.DeleteWithTopic("13")

	category_data, err := category.NewCategoryModel(database.DB).FindOne(context.Background(), 13)

	if err != nil {
		console.Error(err.Error())
	} else {
		fmt.Println(category_data)
		console.Success("删除成功")
	}

}
