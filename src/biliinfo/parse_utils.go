package biliinfo

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func ParseRawInitialState(document string) (*RawInitialState, error) {
	node, err := html.Parse(strings.NewReader(document))
	if err != nil {
		return nil, err
	}
	scripts := getAllScriptTexts(node)
	regx := regexp.MustCompile("window.__INITIAL_STATE__ *= *({.*}) *;")
	var rawInitialState *RawInitialState
	for idx := range scripts {
		script := scripts[idx]
		match := regx.FindStringSubmatch(script)
		if match == nil {
			continue
		}
		json_string := match[1]
		if json_string == "" {
			continue
		}
		rawInitialState = new(RawInitialState)
		err = json.Unmarshal([]byte(json_string), rawInitialState)
		if err != nil {
			log.Println("document unmarshal error")
			continue
		}
		break
	}
	return rawInitialState, nil
}

func ParseRawPlayInfo(document string) (*RawPlayInfo, error) {
	node, err := html.Parse(strings.NewReader(document))
	if err != nil {
		return nil, err
	}
	scripts := getAllScriptTexts(node)
	regx := regexp.MustCompile("window.__playinfo__ *= *({.*}) *")
	var rawInfo *RawPlayInfo
	for idx := range scripts {
		script := scripts[idx]
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
