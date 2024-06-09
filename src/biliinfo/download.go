package biliinfo

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

var session *http.Client = nil

func getClient() *http.Client {
	if session == nil {
		session = new(http.Client)
		session.Jar, _ = cookiejar.New(nil)
	}
	return session
}

func initBiliRequest(u *url.URL) *http.Request {
	req := new(http.Request)
	req.URL = u
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36")
	req.Header.Add("referer", "https://www.bilibili.com")
	return req
}

func resolveAddress(address string) (*url.URL, error) {
	var url_object *url.URL
	if strings.HasPrefix(address, "BV") {
		url_string, err := url.JoinPath("www.bilibili.com/video/", address)
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

func getDocument(req *http.Request) (string, error) {
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
