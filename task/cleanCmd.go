package task

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/dwdcth/cli_sample/rootcmd"
)

// 注册命令
func init() {
	rootcmd.AddCommand(cleanCmd)
}

var cleanCmd = &cobra.Command{
	Short: "clean",
	Use:   "clean ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run clean task command")
	},
}
