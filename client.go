/*
このサービスは、総務省 電波利用ホームページのWeb-API 機能を利用して取得した情報をもとに作成しているが、サービスの内容は総務省によって保証されたものではない
*/

// Package soumu is a simple client to consume radio station etc. information search API of Ministry of Internal Affairs and Communications in Japan.
package soumuradio

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"runtime"
)

// Client is API Client structure
type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
}

// of is indicates API Response to JSON format
const of = "2"
const encoding = "1"

// NewClient is constructer
// 必須の情報が与えられているか、期待するものかをチェックする
func NewClient(cli *http.Client) (*Client, error) {
	urlStr := "https://www.tele.soumu.go.jp"

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURL:    parsedURL,
		HTTPClient: cli,
	}, nil
}

var version = "0.1"
var usrAgent = fmt.Sprintf("SoumuRadioGoClient/%s (%s)", version, runtime.Version())

func (c *Client) newRequest(ctx context.Context, method, spath string, opt requestOptions, body io.Reader) (*http.Request, error) {
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, spath)

	// Add Query parameters to the URL
	params := newParams(opt)
	u.RawQuery = params.Encode() // Escape Query Parameters

	logf("request URL: %v", u.String())

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if ctx == nil {
		ctx = context.Background()
	}
	req = req.WithContext(ctx)

	req.Header.Set("User-Agent", usrAgent)

	// for debug, output proxy setting
	proxy := "no"
	if proxyURL, _ := http.ProxyFromEnvironment(req); proxyURL != nil {
		proxy = proxyURL.String()
	}
	logf("request proxy: %v", proxy)

	return req, nil
}

type requestOptions interface {
	encodeOption() url.Values
}

func newParams(opts requestOptions) url.Values {
	params := opts.encodeOption()
	params.Add("OF", of)
	params.Add("MC", encoding)
	return params
}
