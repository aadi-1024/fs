package main

import (
	"fmt"

	"github.com/aadi-1024/fs/cmds"
	"github.com/alecthomas/kong"
)

func main() {
	ctx := kong.Parse(&cmds.Root{})
	if err := ctx.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
