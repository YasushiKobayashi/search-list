package model

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/YasushiKobayashi/search-list/utils"
	"github.com/pkg/errors"
)

type (
	CompanyInfo struct {
		URL   string
		Tel   string
		Email string
	}
)

func (c *CompanyInfo) GetEmail(doc *goquery.Document) error {
	boby, err := doc.Find("boby").Html()
	if err != nil {
		return errors.Wrap(err, "Html error")
	}

	emails := utils.ExtractionEmail(boby)
	c.Email = strings.Join(emails, "\n")
	return nil
}

func (c *CompanyInfo) GetTel(doc *goquery.Document) error {
	boby, err := doc.Find("boby").Html()
	if err != nil {
		return errors.Wrap(err, "Html error")
	}

	tels := utils.ExtractionTel(boby)
	c.Tel = strings.Join(tels, "\n")
	return nil
}
