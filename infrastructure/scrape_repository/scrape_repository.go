package scrape_repository

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/YasushiKobayashi/search-list/model"
	"github.com/pkg/errors"
)

type (
	Scrape struct{}
)

func (s *Scrape) GetSearchList(k []model.Keyword) (model.SearchLists, error) {
	ch := make(chan bool, 1)
	var err error
	var wg sync.WaitGroup
	var searchLists model.SearchLists
	for _, v := range k {
		wg.Add(1)
		ch <- true
		fmt.Println(v)
		go func(url string) {
			defer func() { <-ch }()
			var searchList *model.SearchList = &model.SearchList{}
			err = getSearchListScrape(searchList, url)
			if err != nil {
				err = errors.Wrap(err, "RunScraipe error")
				log.Println(err)
			}

			searchLists = append(searchLists, searchList)
			time.Sleep(1 * time.Second)
			wg.Done()
		}(v.GetUrl())
	}

	wg.Wait()

	return searchLists, nil
}

func getSearchListScrape(s *model.SearchList, url string) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return errors.Wrap(err, "goquery.NewDocument error")
	}

	err = s.GetListing(doc)
	if err != nil {
		return errors.Wrap(err, "getListing error")
	}

	s.GetSearch(doc)
	return nil
}

func (s *Scrape) GetPageInfoScrage(c *model.CompanyInfo) error {
	doc, err := goquery.NewDocument(c.URL)
	if err != nil {
		return errors.Wrap(err, "goquery.NewDocument error")
	}

	c.GetEmail(doc)
	c.GetTel(doc)
	return nil
}
