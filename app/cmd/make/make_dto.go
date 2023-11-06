package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeDTO = &cobra.Command{
	Use:   "dto",
	Short: " example : make dto  user ",
	Run:   runMakeDTO,
	Args:  cobra.ExactArgs(1),
}

func runMakeDTO(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])

	//  这里面的设计是 每个model 都有三个基本文件, 类似 hook 或者 util 之类的, 所以需要建目录
	dir := fmt.Sprintf("app/dtos/%s_dto/", model.PackageName)

	// 确保父目录和子目录都会创建
	os.MkdirAll(dir, os.ModePerm)

	createFileFromStub(dir+"dto.go", "dto", model)
	createFileFromStub(dir+"mapper.go", "mapper", model)

}
