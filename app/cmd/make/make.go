package make

import (
	"embed"
	"fmt"
	"strings"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/file"
	"github.com/diy0663/gohub/pkg/str"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// plural 复数形式

type Model struct {
	TableName          string
	StructName         string
	StructNamePlural   string
	VariableName       string
	VariableNamePlural string
	PackageName        string
}

//go:embed stubs
var stubsFS embed.FS

var CmdMake = &cobra.Command{
	Use:   "make",
	Short: " Generate file and code ",
}

func init() {
	// 注册子命令
	CmdMake.AddCommand(
		//  自动生成代码类型 cmd ,request ,controller, model 等, 都是 归属 CmdMake 的子命令
		CmdMakeCMD,
		CmdMakeModel,
		CmdMakeAPIController,
		CmdMakeRequest,
	)
}

func makeModelFromString(name string) Model {
	model := Model{}
	// 单数形式的驼峰
	model.StructName = str.Singular(strcase.ToCamel(name))
	model.StructNamePlural = str.Plural(model.StructName)
	// 数据表 下划线复数形式
	model.TableName = str.Snake(model.StructNamePlural)
	// 变量名 小写驼峰形式
	model.VariableName = str.LowerCamel(model.StructName)
	// 包名 下划线形式
	model.PackageName = str.Snake(model.StructName)
	model.VariableNamePlural = str.LowerCamel(model.StructNamePlural)

	return model
}

// 最后一个为可选参数, 假如要传,就得是 map[string]string 类型
func createFileFromStub(filePath string, stubName string, model Model, variables ...interface{}) {

	replaces := make(map[string]string)

	// 把需要额外替换的提前放到替换组里面去
	if len(variables) > 0 {
		replaces = variables[0].(map[string]string)

	}

	if file.Exists(filePath) {
		console.Exit(filePath + " already exists!")
	}
	modelData, err := stubsFS.ReadFile("stubs/" + stubName + ".stub")
	if err != nil {
		console.Exit(err.Error())
	}

	//把读取的模板文件内容变成字符串
	modelStub := string(modelData)
	// 替换
	replaces["{{VariableName}}"] = model.VariableName
	replaces["{{VariableNamePlural}}"] = model.VariableNamePlural
	replaces["{{StructName}}"] = model.StructName
	replaces["{{StructNamePlural}}"] = model.StructNamePlural
	replaces["{{PackageName}}"] = model.PackageName
	replaces["{{TableName}}"] = model.TableName
	// 对模板内容进行替换
	for search, replace := range replaces {
		modelStub = strings.ReplaceAll(modelStub, search, replace)
	}
	// 存放数据到文件去
	err = file.Put([]byte(modelStub), filePath)
	if err != nil {
		console.Exit(err.Error())
	}
	console.Success(fmt.Sprintf("[%s] created. ", filePath))

}
