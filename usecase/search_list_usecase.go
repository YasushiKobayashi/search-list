package usecase

import (
	"github.com/YasushiKobayashi/search-list/model"
	"github.com/pkg/errors"
)

type (
	CsvWriterInteractor struct {
		CsvRepository    CsvRepository
		ScrapeRepository ScrapeRepository
	}

	CsvRepository interface {
		ReadCsv() (*model.CsvWriter, error)
		Write(*model.CsvWriter) error
	}
)

func (i *CsvWriterInteractor) Run() error {
	csvWriter, err := i.readCsv()
	if err != nil {
		return errors.Wrap(err, "ReadCsv error")
	}

	err = i.scrape(csvWriter)
	if err != nil {
		return errors.Wrap(err, "Scrape error")
	}

	return i.writeCsv(csvWriter)
}

func (i *CsvWriterInteractor) readCsv() (*model.CsvWriter, error) {
	csvWriter, err := i.CsvRepository.ReadCsv()
	if err != nil {
		return csvWriter, errors.Wrap(err, "ReadCsv error")
	}

	return csvWriter, nil
}

func (i *CsvWriterInteractor) writeCsv(c *model.CsvWriter) error {
	return i.CsvRepository.Write(c)
}

func (i *CsvWriterInteractor) scrape(c *model.CsvWriter) error {
	searchLists, err := i.ScrapeRepository.Run(c.Keywords)
	if err != nil {
		return errors.Wrap(err, "RunScraipe error")
	}

	c.SearchLists = searchLists
	return nil
}
