package usecase

import (
	"log"
	"sync"

	"github.com/YasushiKobayashi/search-list/model"
	"github.com/YasushiKobayashi/search-list/utils"
	"github.com/pkg/errors"
)

type (
	CompanyInfoInteractor struct {
		CompanyInfoRepository CompanyInfoRepository
		CsvRepository         CsvRepository
	}

	CompanyInfoRepository interface {
		GetPageInfoScrage(*model.CompanyInfo) error
	}
)

func (i *CompanyInfoInteractor) GetPageInfo(parallel int) error {
	csvWriter, err := i.readCsv()
	if err != nil {
		return errors.Wrap(err, "ReadCsv error")
	}

	telCsvWriter, mailCsvWriter := i.getInfo(csvWriter, parallel)
	err = i.CsvRepository.WritePageInfo(telCsvWriter, "tel")
	if err != nil {
		return errors.Wrap(err, "WriteCsv tel error")
	}

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

func (i *CompanyInfoInteractor) readCsv() (*model.CsvWriter, error) {
	csvWriter, err := i.CsvRepository.ReadCsv()
	if err != nil {
		return csvWriter, errors.Wrap(err, "ReadCsv error")
	}

	return csvWriter, nil
}

func (i *CompanyInfoInteractor) getInfo(c *model.CsvWriter, parallel int) (*model.CsvWriter, *model.CsvWriter) {
	var telCsvWriter *model.CsvWriter = &model.CsvWriter{
		Rows: make([][]string, len(c.Rows)),
	}
	var mailCsvWriter *model.CsvWriter = &model.CsvWriter{
		Rows: make([][]string, len(c.Rows)),
	}

	ch := make(chan bool, parallel)
	var wg sync.WaitGroup
	for k, v := range c.Rows {
		wg.Add(1)
		ch <- true
		go func(v []string, k int) {
			defer func() { <-ch }()
			telCsvWriter.Rows[k] = make([]string, len(c.Header))
			mailCsvWriter.Rows[k] = make([]string, len(c.Header))
			for key, val := range v {
				if utils.IsValidUrl(val) {
					var companyInfo *model.CompanyInfo = &model.CompanyInfo{}
					companyInfo.URL = val

					err := i.CompanyInfoRepository.GetPageInfoScrage(companyInfo)
					if err != nil {
						log.Println(err)
					}

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
