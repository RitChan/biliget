package biliinfo

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

var session *http.Client = nil

func GetDocument(address string) (string, error) {
	url, err := resolveAddress(address)
	if err != nil {
		return "", err
	}
	req, err := initBiliGetRequest(url)
	if err != nil {
		return "", err
	}
	client := getClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	text, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(text), nil
}

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

func resolveAddress(address string) (*url.URL, error) {
	var url_object *url.URL
	if strings.HasPrefix(address, "BV") {
		url_string, err := url.JoinPath("https://www.bilibili.com/video/", address)
		if err != nil {
			return nil, err
		}
		u, err := url.Parse(url_string)
		if err != nil {
			return nil, err
		}
		url_object = u
	} else {
		u, err := url.Parse(address)
		if err != nil {
			return nil, err
		}
		url_object = u
	}
	return url_object, nil
}
