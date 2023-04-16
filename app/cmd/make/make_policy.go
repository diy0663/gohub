package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: "Create policy file, example: make policy user",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakePolicy(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])

	os.MkdirAll("app/policies", os.ModePerm)
	// 生成后的目标文件路径
	filePath := fmt.Sprintf("app/policies/%s_policy.go", model.PackageName)

	// 第二个参数就是寻找 stub 目录下的 XX.stub 模板文件
	createFileFromStub(filePath, "policy", model)

}
