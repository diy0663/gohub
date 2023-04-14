package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: " Create request file , example:   make request ",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1),
}

func runMakeRequest(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	filepath := fmt.Sprintf("app/requests/%s_request.go", model.PackageName)
	createFileFromStub(filepath, "request", model)

}
