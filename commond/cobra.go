package commond

import (
	"fmt"
	"github.com/kalifun/gin-template/commond/api"
	"github.com/kalifun/gin-template/utils"
	"github.com/spf13/cobra"
	"os"
)

/*
初始化命令行
*/
var Line = &cobra.Command{
	Use:     "gin-template",
	Short:   "Program main program",
	Long:    "Call gin-template -h for more functions",
	Example: "gin-template -h",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			// 当初没有参数时 提示用户
			tip()
		}
	},
}

func tip() {
	printStr := `You can try using ` + utils.Red("- h") + ` to get more information`
	fmt.Println(printStr)
}

func init() {
	Line.AddCommand(api.Api)
}

func Implement() {
	if err := Line.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
