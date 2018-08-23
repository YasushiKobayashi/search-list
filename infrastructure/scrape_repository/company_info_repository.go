package scrape_repository

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/YasushiKobayashi/search-list/model"
	"github.com/YasushiKobayashi/search-list/utils"
	"github.com/pkg/errors"
)

type (
	CompanyInfoRepository struct{}
)

func (r *CompanyInfoRepository) GetPageInfoScrage(c *model.CompanyInfo) error {
	doc, err := newDocument(c.URL)
	if err != nil {
		return errors.Wrap(err, "goquery.NewDocument error")
	}

	err = r.getEmail(doc, c)
	if err != nil {
		return errors.Wrap(err, "getEmail error")
	}

	err = r.getTel(doc, c)
	if err != nil {
		return errors.Wrap(err, "getTel error")
	}
	return nil
}

func (r *CompanyInfoRepository) getEmail(doc *goquery.Document, c *model.CompanyInfo) error {
	boby, err := doc.Html()
	if err != nil {
		return errors.Wrap(err, "Html error")
	}

	const mailTUrlPrefix = "mailto:"
	var emails []string
	doc.Each(func(_ int, s *goquery.Selection) {
		urlStr, exist := s.Attr("href")
		if exist && strings.HasPrefix(urlStr, mailTUrlPrefix) {
			email := strings.Replace(urlStr, mailTUrlPrefix, "", -1)
			emails = append(emails, email)
		}
	})

	extractionEmails := utils.ExtractionEmail(boby)
	emails = append(emails, extractionEmails...)
	list := utils.UniqStringArray(emails)
	c.Email = strings.Join(list, "\n")
	return nil
}

func (r *CompanyInfoRepository) getTel(doc *goquery.Document, c *model.CompanyInfo) error {
	boby, err := doc.Html()
	if err != nil {
		return errors.Wrap(err, "Html error")
	}

	const telTUrlPrefix = "tel:"
	var tels []string
	doc.Each(func(_ int, s *goquery.Selection) {
		urlStr, exist := s.Attr("href")
		if exist && strings.HasPrefix(urlStr, telTUrlPrefix) {
			tel := strings.Replace(urlStr, telTUrlPrefix, "", -1)
			tels = append(tels, tel)
		}
	})

	extractionTels := utils.ExtractionTel(boby)
	tels = append(tels, extractionTels...)
	list := utils.UniqStringArray(tels)
	c.Tel = strings.Join(list, "\n")
	return nil
}
