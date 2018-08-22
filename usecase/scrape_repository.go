package usecase

import (
	"github.com/YasushiKobayashi/search-list/model"
)

type (
	ScrapeRepository interface {
		Run([]model.Keyword) (model.SearchLists, error)
	}
)
