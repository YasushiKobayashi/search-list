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
	CompanyInfoHandler struct {
		Interactor usecase.CompanyInfoInteractor
	}
)

func NewCompanyInfoHandler(path string) *CompanyInfoHandler {
	return &CompanyInfoHandler{
		Interactor: usecase.CompanyInfoInteractor{
			CompanyInfoRepository: &scrape_repository.CompanyInfoRepository{},
			CsvRepository: &csv_repository.CsvRepository{
				Path: path,
			},
		},
	}
}

var getPageInfo = cli.Command{
	Name:    "page",
	Aliases: []string{"p"},
	Usage:   "...",
	Description: `
Get scraped pape company tel and email.
`,
	Action: getPageInfoHandler,
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

func getPageInfoHandler(c *cli.Context) {
	filPath := c.String("file")
	parallel := c.String("parallel")
	parallelNumber, err := strconv.Atoi(parallel)
	if err != nil {
		err = errors.Wrap(err, "parallel must number.")
		panic(err)
	}

	handler := NewCompanyInfoHandler(filPath)
	err = handler.Interactor.GetPageInfo(parallelNumber)
	if err != nil {
		panic(err)
	}
}
