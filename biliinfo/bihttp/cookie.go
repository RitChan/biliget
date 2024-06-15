package bihttp

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

func ClientAddCookies(client *http.Client, cookies []*http.Cookie) {
	if client.Jar == nil {
		return
	}
	u, _ := url.Parse(urlBilibili)
	oldCookies := client.Jar.Cookies(u)
	client.Jar.SetCookies(u, append(oldCookies, cookies...))
}

func DumpCookies(filepath string, cookies []*http.Cookie) error {
	bytes, err := json.Marshal(cookies)
	if err != nil {
		return err
	}
	os.WriteFile(filepath, bytes, 0666)
	log.Printf("Write cookie file: %s\n", filepath)
	return nil
}

func LoadCookies(filepath string) ([]*http.Cookie, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var cookies []*http.Cookie
	err = json.Unmarshal(bytes, &cookies)
	if err != nil {
		return nil, err
	}
	return cookies, nil
}

func parseRawCookies(rawCookies string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", rawCookies)
	req := http.Request{Header: header}
	return req.Cookies()
}
