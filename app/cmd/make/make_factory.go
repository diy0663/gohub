package make

import (
	"fmt"
	"os"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/spf13/cobra"
)

var CmdMakeFactory = &cobra.Command{
	Use:   "factory",
	Short: "Create model's factory file, example: make factory user",
	Args:  cobra.ExactArgs(1),
	Run:   runMakeFactory,
}

func runMakeFactory(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	// 最好检查那个model 是否已存在,位置是相对于 main.go 而言
	_, err := os.Stat("app/models/" + model.PackageName)
	if os.IsNotExist(err) {
		console.Exit("需要先生成" + model.PackageName + "对应的model")
	}

	filePath := fmt.Sprintf("database/factories/%v_factory.go", model.PackageName)
	createFileFromStub(filePath, "factory", model)

}
