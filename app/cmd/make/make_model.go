package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: " example : make model model article ",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1),
}

func runMakeModel(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])

	//  这里面的设计是 每个model 都有三个基本文件, 类似 hook 或者 util 之类的, 所以需要建目录
	dir := fmt.Sprintf("app/models/%s/", model.PackageName)

	// 确保父目录和子目录都会创建
	os.MkdirAll(dir, os.ModePerm)

	createFileFromStub(dir+model.PackageName+"_model.go", "model/model", model)
	createFileFromStub(dir+model.PackageName+"_util.go", "model/model_util", model)
	createFileFromStub(dir+model.PackageName+"_hooks.go", "model/model_hooks", model)

}
