package biliinfo

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func ParseBvInfo(document string) (*BvInfo, error) {
	initialState, err := ParseRawInitialState(document)
	if err != nil {
		return nil, err
	}
	playInfo, err := ParseRawPlayInfo(document)
	if err != nil {
		return nil, err
	}
	bvinfo := new(BvInfo)
	// Basic
	bvinfo.BVID = initialState.BVID
	bvinfo.Title = initialState.VideoData.Title
	bvinfo.Duration = float32(playInfo.Data.Dash.Duration)
	bvinfo.PartID = int32(initialState.P)
	// Up
	bvinfo.Up.Name = initialState.UpData.Name
	bvinfo.Up.Sex = initialState.UpData.Sex
	bvinfo.Up.Profile = initialState.UpData.ProfileURL
	// Videos
	for _, rawVideo := range playInfo.Data.Dash.Video {
		videoInfo := new(videoInfo)
		videoInfo.Url = rawVideo.BaseUrl
		videoInfo.UrlBackup = rawVideo.BackUpUrl
		videoInfo.MimeType = rawVideo.MimeType
		videoInfo.Width = uint32(rawVideo.Width)
		videoInfo.Height = uint32(rawVideo.Heihgt)
		frameRate, err := strconv.ParseFloat(rawVideo.FrameRate, 32)
		if err != nil {
			log.Printf("Warning: cannot parse frame rate: %s", rawVideo.FrameRate)
		} else {
			videoInfo.FrameRate = float32(frameRate)
		}
		bvinfo.Videos = append(bvinfo.Videos, *videoInfo)
	}
	// Audios
	for _, rawAudio := range playInfo.Data.Dash.Audio {
		audioInfo := new(audioInfo)
		audioInfo.Url = rawAudio.BaseUrl
		audioInfo.UrlBackup = rawAudio.BackUpUrl
		audioInfo.MimeType = rawAudio.MimeType
		bvinfo.Audios = append(bvinfo.Audios, *audioInfo)
	}
	// Parts
	for _, rawPage := range initialState.VideoData.Pages {
		part := partInfo{}
		part.PartId = int32(rawPage.PageID)
		part.PartName = rawPage.PageTitle
		bvinfo.Parts = append(bvinfo.Parts, part)
	}
	// Statistics
	bvinfo.Statistics.View = uint32(initialState.VideoData.Stat.View)
	bvinfo.Statistics.Like = uint32(initialState.VideoData.Stat.Like)
	bvinfo.Statistics.Coin = uint32(initialState.VideoData.Stat.Coin)
	bvinfo.Statistics.Favourite = uint32(initialState.VideoData.Stat.Favourite)
	bvinfo.Statistics.Reply = uint32(initialState.VideoData.Stat.Reply)

	return bvinfo, nil
}

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
