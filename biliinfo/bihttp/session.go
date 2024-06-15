package bihttp

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

var session *http.Client = nil

func getClient() *http.Client {
	if session == nil {
		session = new(http.Client)
		session.Jar, _ = cookiejar.New(nil)
	}
	return session
}

func initBiliGetRequest(u *url.URL) (*http.Request, error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.URL = u
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36")
	req.Header.Set("referer", "https://www.bilibili.com")
	return req, nil
}
