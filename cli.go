package main

import (
	"fmt"
	"os"
	"github.com/dwdcth/cli_sample/rootcmd"
	_"github.com/dwdcth/cli_sample/task"
	_"github.com/dwdcth/cli_sample/update"
)



func main() {

	if err := rootcmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
