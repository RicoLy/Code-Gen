package main

import (
	"code-gen/core"
	"github.com/urfave/cli/v2"
)

var (
	stop chan bool // 关闭信号
	app  *cli.App  // cli对象
)

func main() {
	app = cli.NewApp()
	stop = make(chan bool)
	core.Start(app, stop)
}
