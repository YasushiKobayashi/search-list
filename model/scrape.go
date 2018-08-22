package model

import "net/url"

type (
	Keyword string
)

func (k Keyword) GetUrl() string {
	const baseUrl string = "https://www.google.co.jp/search?"
	v := url.Values{}
	v.Set("q", k.String())
	return baseUrl + v.Encode()
}

func (k Keyword) String() string {
	return string(k)
}
