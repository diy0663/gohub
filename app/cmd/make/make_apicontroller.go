package make

import (
	"fmt"
	"os"
	"strings"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/spf13/cobra"
)

var CmdMakeAPIController = &cobra.Command{
	Use: "apicontroller",
	// 虽然后面只接收1个参数, 但是这个参数其实是2个, 这里面带了版本号
	Short: " example:  make apicontroller v1/project ",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1),
}

func runMakeAPIController(cmd *cobra.Command, args []string) {

	parameter := strings.Split(args[0], "/")
	if len(parameter) != 2 {
		console.Exit(" apicontroller format: v1/project ,should with / ")
	}

	apiVersion, modelName := parameter[0], parameter[1]
	model := makeModelFromString(modelName)

	// 确保目录创建成功
	os.MkdirAll("app/http/controllers/"+apiVersion+"/", os.ModePerm)

	filePath := fmt.Sprintf("app/http/controllers/%v/%v_controller.go", apiVersion, model.TableName)

	createFileFromStub(filePath, apiVersion+"_apicontroller", model, map[string]string{"{{APIVersion}}": apiVersion})
}
