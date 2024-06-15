package biliinfo

import (
	"encoding/json"
	"net/url"
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
