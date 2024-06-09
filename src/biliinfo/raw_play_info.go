package biliinfo

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type RawPlayInfo struct {
	Data RawPlayInfoData `json:"data"`
}

type RawPlayInfoData struct {
	Dash RawPlayInfoDataDash `json:"dash"`
}

type RawPlayInfoDataDash struct {
	Duration int32                      `json:"duration"` // seconds
	Video    []RawPlayInfoDataDashVideo `json:"video"`
	Audio    []RawPlayInfoDataDashAudio `json:"audio"`
}

type RawPlayInfoDataDashVideo struct {
	BaseUrl   string   `json:"baseUrl"`
	BackUpUrl []string `json:"backupUrl"`
	MimeType  string   `json:"mimeType"`
	Width     int32    `json:"width"`
	Heihgt    int32    `json:"height"`
	FrameRate string   `json:"frameRate"`
}

type RawPlayInfoDataDashAudio struct {
	BaseUrl   string   `json:"baseUrl"`
	BackUpUrl []string `json:"backupUrl"`
	MimeType  string   `json:"mimeType"`
}

func ParseRawPlayInfo(document string) (*RawPlayInfo, error) {
	node, err := html.Parse(strings.NewReader(document))
	if err != nil {
		return nil, err
	}
	scripts := getAllScriptTexts(node)
	var rawInfo *RawPlayInfo
	for idx := range scripts {
		script := scripts[idx]
		regx := regexp.MustCompile("window.__playinfo__ *= *({.*}) *")
		match := regx.FindStringSubmatch(script)
		if match == nil {
			continue
		}
		json_string := match[1]
		if json_string == "" {
			continue
		}
		rawInfo = new(RawPlayInfo)
		err = json.Unmarshal([]byte(json_string), rawInfo)
		if err != nil {
			log.Println("document unmarshal error")
			continue
		}
		break
	}
	return rawInfo, nil
}

func getAllScriptTexts(node *html.Node) []string {
	result := make([]string, 0, 5)
	if node.Type == html.ElementNode && node.Data == "script" {
		for child := node.FirstChild; child != nil && child.Type == html.TextNode; child = child.NextSibling {
			result = append(result, child.Data)
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		result = append(result, getAllScriptTexts(child)...)
	}
	return result
}
