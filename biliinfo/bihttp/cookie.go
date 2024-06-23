package bihttp

import (
	"net/http"
	"net/url"
)

func SetCookies(cookies []*http.Cookie) {
	ClientSetCookies(GetClient(), cookies)
}

func ClientSetCookies(client *http.Client, cookies []*http.Cookie) {
	if client.Jar == nil {
		return
	}
	u, _ := url.Parse(urlBilibili)
	oldCookies := client.Jar.Cookies(u)
	mapping := make(map[string]*http.Cookie) // name -> cookie
	for _, c := range oldCookies {
		mapping[c.Name] = c
	}
	for _, c := range cookies {
		mapping[c.Name] = c
	}
	newCookies := make([]*http.Cookie, 0, len(mapping))
	for _, c := range mapping {
		newCookies = append(newCookies, c)
	}
	client.Jar.SetCookies(u, newCookies)
}

func parseRawCookies(rawCookies string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", rawCookies)
	req := http.Request{Header: header}
	return req.Cookies()
}
