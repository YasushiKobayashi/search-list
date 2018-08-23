package scrape_repository

import (
	"github.com/YasushiKobayashi/search-list/model"
	"github.com/pkg/errors"
)

type (
	SearchListRepository struct{}
)

func (r *SearchListRepository) GetSearchList(s *model.SearchList, urlStr string) error {
	doc, err := newDocument(urlStr)
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
