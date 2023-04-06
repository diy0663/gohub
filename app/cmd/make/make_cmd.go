package make

import (
	"fmt"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/spf13/cobra"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, should be snake_case, exmaple: make cmd buckup_database",
	Run:   runMakeCMD,
	// 只允许传1个参数
	Args: cobra.ExactArgs(1),
}

func runMakeCMD(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])
	filePath := fmt.Sprintf("app/cmd/%s.go", model.PackageName)
	createFileFromStub(filePath, "cmd", model)
	// 友好提示
	console.Success("command name:" + model.PackageName)
	console.Success("command variable name: cmd.Cmd" + model.StructName)
	console.Warning("please resave the file to auto complete the import codes ")
	console.Warning("please edit main.go's app.Commands slice to register command")
}
