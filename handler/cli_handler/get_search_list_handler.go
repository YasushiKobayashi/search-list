package cli_handler

import (
	"strconv"

	"github.com/YasushiKobayashi/search-list/infrastructure/csv_repository"
	"github.com/YasushiKobayashi/search-list/infrastructure/scrape_repository"
	"github.com/YasushiKobayashi/search-list/usecase"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

type (
	CsvWriterHandler struct {
		Interactor usecase.CsvWriterInteractor
	}
)

func NewCsvWriterHandler(path string) *CsvWriterHandler {
	return &CsvWriterHandler{
		Interactor: usecase.CsvWriterInteractor{
			CsvRepository: &csv_repository.CsvRepository{
				Path: path,
			},
			SearchListRepository: &scrape_repository.SearchListRepository{},
		},
	}
}

var getSearchList = cli.Command{
	Name:    "search-list",
	Aliases: []string{"s"},
	Usage:   "...",
	Description: `
Get google top list.
`,
	Action: getSearchListHandler,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Value: "sample.csv",
			Usage: "csv file path",
		},
		cli.StringFlag{
			Name:  "parallel, p",
			Value: "1",
			Usage: "palallel number goroutine",
		},
	},
}

func getSearchListHandler(c *cli.Context) {
	filePath := c.String("file")
	parallel := c.String("parallel")
	parallelNumber, err := strconv.Atoi(parallel)
	if err != nil {
		err = errors.Wrap(err, "parallel must number.")
		panic(err)
	}

	handler := NewCsvWriterHandler(filePath)
	err = handler.Interactor.GetSearchList(parallelNumber)
	if err != nil {
		panic(err)
	}
}
