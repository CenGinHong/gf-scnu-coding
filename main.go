package main

import (
	_ "gf-scnu-coding/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"gf-scnu-coding/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
