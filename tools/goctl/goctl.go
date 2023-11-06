package main

import (
	"github.com/userzhangjinlong/go-zero/core/load"
	"github.com/userzhangjinlong/go-zero/core/logx"
	"github.com/userzhangjinlong/go-zero/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
