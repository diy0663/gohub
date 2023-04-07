package make

import (
	"fmt"
	"os"
	"strings"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/spf13/cobra"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller，exmaple: make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1),
}

func runMakeAPIController(cmd *cobra.Command, args []string) {
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/user")
	}
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	filePath := fmt.Sprintf("app/http/controllers/api/%s/%s_controller.go", apiVersion, model.TableName)
	dir := fmt.Sprintf("app/http/controllers/api/%s/", apiVersion)
	// 确保目录会被创建
	os.MkdirAll(dir, os.ModePerm)

	// 基于模板创建文件（做好变量替换）
	// 需要把api版本也穿进去
	extra_replace := map[string]string{"{{API_VERSION}}": apiVersion}

	createFileFromStub(filePath, "apicontroller", model, extra_replace)

}
