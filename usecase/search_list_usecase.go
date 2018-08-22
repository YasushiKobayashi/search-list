package usecase

import (
	"fmt"
	"log"
	"sync"

	"github.com/YasushiKobayashi/search-list/model"
	"github.com/YasushiKobayashi/search-list/utils"
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
		WritePageInfo(*model.CsvWriter, string) error
	}
)

func (i *CsvWriterInteractor) GetSearchList() error {
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
	searchLists, err := i.ScrapeRepository.GetSearchList(c.Keywords)
	if err != nil {
		return errors.Wrap(err, "RunScraipe error")
	}

	c.SearchLists = searchLists
	return nil
}

func (i *CsvWriterInteractor) GetPageInfo(parallel int) error {
	csvWriter, err := i.readCsv()
	if err != nil {
		return errors.Wrap(err, "ReadCsv error")
	}

	telCsvWriter, mailCsvWriter := i.getInfo(csvWriter, parallel)
	err = i.CsvRepository.WritePageInfo(telCsvWriter, "tel")
	if err != nil {
		return errors.Wrap(err, "WriteCsv tel error")
	}

	err = i.CsvRepository.WritePageInfo(mailCsvWriter, "mail")
	if err != nil {
		return errors.Wrap(err, "WriteCsv mail error")
	}
	return nil
}

func (i *CsvWriterInteractor) getInfo(c *model.CsvWriter, parallel int) (*model.CsvWriter, *model.CsvWriter) {
	telCsvWriter := c
	mailCsvWriter := c

	ch := make(chan bool, parallel)
	var wg sync.WaitGroup
	for k, v := range c.Rows {
		wg.Add(1)
		ch <- true

		go func(v []string, k int) {
			for key, val := range v {
				if utils.IsValidUrl(val) {
					var companyInfo *model.CompanyInfo = &model.CompanyInfo{}
					companyInfo.URL = val

					fmt.Println(companyInfo)
					err := i.ScrapeRepository.GetPageInfoScrage(companyInfo)
					if err != nil {
						log.Println(err)
					}
					fmt.Println(companyInfo)

					telCsvWriter.Rows[k][key] = companyInfo.Tel
					mailCsvWriter.Rows[k][key] = companyInfo.Email
				} else {
					telCsvWriter.Rows[k][key] = val
					mailCsvWriter.Rows[k][key] = val
				}
			}
			wg.Done()
		}(v, k)
	}

	wg.Wait()
	return telCsvWriter, mailCsvWriter
}
