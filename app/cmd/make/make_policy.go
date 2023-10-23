package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// 授权策略判断文件的生成
var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: "example: make policy user",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1),
}

func runMakePolicy(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	os.MkdirAll("app/policies", os.ModePerm)
	filePath := fmt.Sprintf("app/policies/%s_policy.go", model.PackageName)
	createFileFromStub(filePath, "policy", model)

}
