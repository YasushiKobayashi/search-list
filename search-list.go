package main

import (
	"os"

	"github.com/YasushiKobayashi/search-list/handler/cli_handler"
	"github.com/urfave/cli"
)

var Version string = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "search-list"
	app.Usage = "get google search top & listing list"
	app.Author = "Yasushi Kobayashi"
	app.Email = "ptpadan@gmail.com"
	app.Version = Version
	app.Commands = cli_handler.Commands

	app.Run(os.Args)
}
