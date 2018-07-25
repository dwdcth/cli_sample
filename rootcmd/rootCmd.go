package rootcmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OK")
	},
}
//封装了两个函数
func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func Execute() error {
	return rootCmd.Execute()
}