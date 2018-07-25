package update

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/dwdcth/cli_sample/rootcmd"
)

// 注册命令
func init() {
	rootcmd.AddCommand(userCmd)
}

var userCmd = &cobra.Command{
	Short: "user",
	Use:   "user ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run user update command")
	},
}
