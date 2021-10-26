package main

import (
	"fmt"
	"os"

	"robot/internal/command"
	"robot/internal/table"
)

func main() {
	params := os.Args[1:]
	if len(params) == 0 {
		fmt.Printf("missing file name from the argument list\n")
		fmt.Printf("expected usage: ./robot commands.txt\n")
		os.Exit(1)
	}
	fileName := params[0]

	tbl := table.New(5, 5)
	cmds, err := command.ScanCommandList(fileName)
	if err != nil {
		fmt.Printf("failed to scan command list: %s\n", err.Error())
		os.Exit(1)
	}

	for _, cmd := range cmds {
		cmd(tbl)
	}
}
