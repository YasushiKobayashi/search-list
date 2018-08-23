package utils

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func GetRedilectedUrl(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	cl := http.Client{}
	var lastUrl string
	cl.CheckRedirect = func(req *http.Request, via []*http.Request) error {

		if len(via) > 10 {
			return errors.New("too many redirects")
		}
		lastUrl = req.URL.String()
		return nil
	}
	_, err = cl.Do(req)
	if err != nil {
		return lastUrl, errors.Wrap(err, "cl.Do error")
	}

	return lastUrl, nil
}

func ParseGoogleUrl(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", errors.Wrap(err, "url.Parse error")
	}
	q := u.Query()
	if q != nil && q["q"] != nil {
		return q["q"][0], nil
	}
	return "", fmt.Errorf("invalid url : %s", urlStr)
}
