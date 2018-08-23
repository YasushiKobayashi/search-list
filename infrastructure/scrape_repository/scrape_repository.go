package scrape_repository

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YasushiKobayashi/search-list/utils"
	"github.com/pkg/errors"
)

func newDocument(urlStr string) (res *goquery.Document, err error) {
	req, err := utils.RequestGet(urlStr)
	if err != nil {
		return res, errors.Wrap(err, "RequestGet error")
	}

	document, err := goquery.NewDocumentFromReader(req.Body)
	if err != nil {
		return res, errors.Wrap(err, "NewDocumentFromReader error")
	}

	return document, nil
}
