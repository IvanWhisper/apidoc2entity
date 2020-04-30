package app

import (
	"apidoc2entity/cmd"
	"fmt"
)

func CliAppRun() {
	err := cmd.Run()
	if err != nil {
		fmt.Print(err)
	}
}
