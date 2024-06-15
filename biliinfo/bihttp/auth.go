package bihttp

import (
	"encoding/json"
	"io"
	"net/url"
)

type getQrcodeUrlResp struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	TTL     int                  `json:"ttl"`
	Data    getQrcodeUrlRespData `json:"data"`
}

type getQrcodeUrlRespData struct {
	Url       string `json:"url"`
	QrcodeKey string `json:"qrcode_key"`
}

type qrCodePollResp struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    qrCodePollRespData `json:"data"`
}

type qrCodePollRespData struct {
	Url          string `json:"url"`
	RefreshToken string `json:"refresh_token"`
	Timestamp    string `json:"timestamp"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
}

func BiliGetQrcodeUrl() (*getQrcodeUrlResp, error) {
	resp := new(getQrcodeUrlResp)
	err := BiliGetUrlJson(urlQrcodeGenerate, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Returns (resp json struct, Set-Cookie header values, error)
func BiliQrcodePoll(qrcodeKey string) (*qrCodePollResp, []string, error) {
	// Construct request
	u, err := url.Parse(urlQrcodePoll)
	if err != nil {
		return nil, nil, err
	}
	req, err := initBiliGetRequest(u)
	if err != nil {
		return nil, nil, err
	}
	q := req.URL.Query()
	q.Set("qrcode_key", qrcodeKey)
	req.URL.RawQuery = q.Encode()

	// Send
	resp, err := biliDoRequest(req)
	if err != nil {
		return nil, nil, err
	}

	// Parse
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	jsonResult := new(qrCodePollResp)
	err = json.Unmarshal(bytes, jsonResult)
	if err != nil {
		return nil, nil, err
	}
	setCookies := resp.Header.Values("Set-Cookie")
	return jsonResult, setCookies, err
}
