package api

import "github.com/spf13/cobra"

var Api = &cobra.Command{
	Use:     "server",
	Short:   "Start Server",
	Long:    "Start Api Server",
	Example: "gin-template server",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func RunServer() {

}
