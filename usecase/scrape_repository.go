package usecase

import (
	"github.com/YasushiKobayashi/search-list/model"
)

type (
	ScrapeRepository interface {
		GetSearchList([]model.Keyword) (model.SearchLists, error)
		GetPageInfoScrage(*model.CompanyInfo) error
	}
)
