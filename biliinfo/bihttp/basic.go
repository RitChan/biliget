package bihttp

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func BiliGetAddress(address string) ([]byte, error) {
	u, err := resolveAddress(address)
	if err != nil {
		return nil, err
	}
	bytes, err := biliGetUrlObject(u)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func BiliGetUrl(url_string string) ([]byte, error) {
	u, err := url.Parse(url_string)
	if err != nil {
		return nil, err
	}
	bytes, err := biliGetUrlObject(u)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// "v" will be parsed into json.Unmarshal so parsing result will be stored in "v"
func BiliGetUrlJson(url_string string, v any) error {
	bytes, err := BiliGetUrl(url_string)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return err
	}
	return nil
}

func biliGetUrlObject(u *url.URL) ([]byte, error) {
	req, err := initBiliGetRequest(u)
	if err != nil {
		return nil, err
	}
	resp, err := biliDoRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func biliDoRequest(req *http.Request) (*http.Response, error) {
	client := getClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
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
