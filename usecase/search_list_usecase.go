package usecase

import (
	"fmt"
	"log"
	"sync"

	"github.com/YasushiKobayashi/search-list/model"
	"github.com/pkg/errors"
)

type (
	CsvWriterInteractor struct {
		CsvRepository        CsvRepository
		SearchListRepository SearchListRepository
	}

	SearchListRepository interface {
		GetSearchList(*model.SearchList, string) error
	}

	CsvRepository interface {
		ReadCsv() (*model.CsvWriter, error)
		WriteSearchList(*model.CsvWriter) error
		WritePageInfo(*model.CsvWriter, string) error
	}
)

func (i *CsvWriterInteractor) GetSearchList(palallel int) error {
	csvWriter, err := i.readCsv()
	if err != nil {
		return errors.Wrap(err, "ReadCsv error")
	}

	err = i.getSearchList(csvWriter, palallel)
	if err != nil {
		return errors.Wrap(err, "Scrape error")
	}

	return i.writeCsv(csvWriter)
}

func (i *CsvWriterInteractor) getSearchList(c *model.CsvWriter, palallel int) error {
	var searchLists model.SearchLists
	var err error

	ch := make(chan bool, palallel)
	var wg sync.WaitGroup
	for _, v := range c.Keywords {
		wg.Add(1)
		ch <- true
		fmt.Println(v)
		go func(url string) {
			defer func() { <-ch }()
			var searchList *model.SearchList = &model.SearchList{}
			err = i.SearchListRepository.GetSearchList(searchList, url)
			if err != nil {
				err = errors.Wrap(err, "RunScraipe error")
				log.Println(err)
			}

			searchLists = append(searchLists, searchList)
			wg.Done()
		}(v.GetUrl())
	}

	wg.Wait()

	c.SearchLists = searchLists
	return nil
}

func (i *CsvWriterInteractor) readCsv() (*model.CsvWriter, error) {
	csvWriter, err := i.CsvRepository.ReadCsv()
	if err != nil {
		return csvWriter, errors.Wrap(err, "ReadCsv error")
	}

	return csvWriter, nil
}

func (i *CsvWriterInteractor) writeCsv(c *model.CsvWriter) error {
	return i.CsvRepository.WriteSearchList(c)
}
