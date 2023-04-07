package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Crate model file, example: make model user",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1),
}

func runMakeModel(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])

	dir := fmt.Sprintf("app/models/%s/", model.PackageName)
	// 确保目录会被创建
	os.MkdirAll(dir, os.ModePerm)

	// 替换变量 ,生成对应的3个model相关的文件

	// xx_model 里面放对应的结构体以及基本的增删改查 ,基于结构体的方法
	createFileFromStub(dir+model.PackageName+"_model.go", "model/model", model)
	// XX_util里面放跟数据表相关的查询函数 (不是基于结构体的方法)
	createFileFromStub(dir+model.PackageName+"_util.go", "model/model_util", model)
	// XX_hooks 里面放一些  beforesave 之类的
	createFileFromStub(dir+model.PackageName+"_hooks.go", "model/model_hooks", model)

}
