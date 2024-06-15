package bihttp

import (
	"net/http"
	"net/url"
)

func ClientAddCookies(client *http.Client, cookies []*http.Cookie) {
	if client.Jar == nil {
		return
	}
	u, _ := url.Parse(urlBilibili)
	oldCookies := client.Jar.Cookies(u)
	client.Jar.SetCookies(u, append(oldCookies, cookies...))
}

func parseRawCookies(rawCookies string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", rawCookies)
	req := http.Request{Header: header}
	return req.Cookies()
}
