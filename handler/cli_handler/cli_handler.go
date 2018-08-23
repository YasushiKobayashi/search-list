package cli_handler

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	getSearchList,
	getPageInfo,
}
