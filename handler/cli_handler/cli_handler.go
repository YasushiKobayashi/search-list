package cli_handler

import (
	"github.com/YasushiKobayashi/search-list/infrastructure/csv_repository"
	"github.com/YasushiKobayashi/search-list/infrastructure/scrape"
	"github.com/YasushiKobayashi/search-list/usecase"
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
			ScrapeRepository: &scrape.Scrape{},
		},
	}
}

var Commands = []cli.Command{
	commandGetSearchList,
}
var commandGetSearchList = cli.Command{
	Name:    "get",
	Aliases: []string{"g"},
	Usage:   "...",
	Description: `
upload dump file to S3
`,
	Action: GetSearchList,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Value: "sample.csv",
			Usage: "csv file path",
		},
		// cli.StringFlag{
		// 	Name:  "parallel, p",
		// 	Value: "palallel number goroutine",
		// 	Usage: "",
		// },
	},
}

func GetSearchList(c *cli.Context) error {
	path := c.String("file")
	handler := NewCsvWriterHandler(path)
	handler.Interactor.Run()
	return nil
}
