package model

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/YasushiKobayashi/search-list/utils"
	"github.com/pkg/errors"
)

type (
	SearchLists []*SearchList

	SearchList struct {
		Listing1 string
		Listing2 string
		Listing3 string
		Listing4 string
		Search1  string
		Search2  string
		Search3  string
		Search4  string
	}
)

func (s *SearchList) GetListing(doc *goquery.Document) error {
	var urls []string
	doc.Find("div#center_col h3.ellip a").Each(func(_ int, s *goquery.Selection) {
		url, exist := s.Attr("href")
		if exist {
			urls = append(urls, url)
		}
	})

	for i, v := range urls {
		if v != "" {
			err := s.SetListing(v, i)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}

func (s *SearchList) SetListing(url string, number int) (err error) {
	url, err = utils.GetRedilectedUrl(url)
	if err != nil {
		return errors.Wrap(err, "GetRedilectedUrl error")
	}

	switch {
	case number == 0:
		s.Listing1 = url
	case number == 1:
		s.Listing2 = url
	case number == 2:
		s.Listing3 = url
	case number == 3:
		s.Listing4 = url
	}
	return nil
}

func (s *SearchList) GetSearch(doc *goquery.Document) error {
	var urls []string
	doc.Find("div#ires h3.r a").Each(func(_ int, s *goquery.Selection) {
		href, exist := s.Attr("href")
		if exist {
			urlStr, _ := utils.ParseGoogleUrl(href)
			urls = append(urls, urlStr)
		}
	})

	for i, v := range urls {
		s.SetSearch(v, i)
	}

	return nil
}

func (s *SearchList) SetSearch(url string, number int) {
	switch number {
	case 0:
		s.Search1 = url
	case 1:
		s.Search2 = url
	case 2:
		s.Search3 = url
	case 3:
		s.Search4 = url
	}
}

func (s *SearchList) GetCsvVal() (res []string) {
	if s == nil {
		return res
	}
	return []string{s.Listing1, s.Listing2, s.Listing3, s.Listing4, s.Search1, s.Search2, s.Search3, s.Search4}
}
