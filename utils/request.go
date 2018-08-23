package utils

import (
	"net/http"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/proxy"
)

// RequestGetByTor
// http request get, proxy by tor
// refs: https://gist.github.com/mmcloughlin/17e3ca302785f0e525655191d3f9211d
func RequestGetByTor(urlStr string) (res *http.Response, err error) {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9150", nil, proxy.Direct)
	if err != nil {
		return res, errors.Wrap(err, "proxy.SOCKS5 error")
	}

	// Setup HTTP transport
	tr := &http.Transport{
		Dial: dialer.Dial,
	}
	client := &http.Client{Transport: tr}

	res, err = client.Get(urlStr)
	if err != nil {
		return res, errors.Wrap(err, "client.Get error")
	}
	return res, nil
}

func RequestGet(urlStr string) (res *http.Response, err error) {
	timeout := time.Duration(15 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	res, err = client.Get(urlStr)
	if err != nil {
		return res, errors.Wrap(err, "client.Get error")
	}
	return res, nil
}
