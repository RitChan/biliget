package main

import (
	"biliget/biliinfo/bihttp"
	"biliget/biliinfo/bvinfo"
	"biliget/config"
	"biliget/ffmpeg"
	"biliget/gui"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	config.Initialize()
	// gui.RunGiuExample()
	gui.RunGiuMain()
}

func Debug() {
	bytes, err := bihttp.BiliGetAddress("BV1YJ4m1g7EQ")
	check(err)
	err = os.WriteFile(".temp/debug.html", bytes, 0666)
	check(err)
	info, err := bvinfo.ParseBvInfo(string(bytes))
	check(err)
	videoPath := ".temp/video.mp4"
	if len(info.Videos) > 0 {
		bytes, err := bihttp.BiliGetUrl(info.Videos[0].UrlBackup[0])
		check(err)
		err = os.WriteFile(videoPath, bytes, 0666)
		check(err)
		log.Printf("Download video (%d, %d)", info.Videos[0].Width, info.Videos[0].Height)
	}
	audioPath := ".temp/audio.m4a"
	if len(info.Audios) > 0 {
		bytes, err := bihttp.BiliGetUrl(info.Audios[0].UrlBackup[0])
		check(err)
		err = os.WriteFile(audioPath, bytes, 0666)
		check(err)
	}
	err = ffmpeg.MergeAudioVideo(audioPath, videoPath, ".temp/merged.mp4")
	check(err)
	err = ffmpeg.ConvertAudio(audioPath, strings.TrimSuffix(audioPath, filepath.Ext(audioPath))+".mp3")
	check(err)
	// client := bihttp.GetClient()
	// u, _ := url.Parse("https://www.bilibili.com")
	// bihttp.DumpCookies(".temp/cookie.json", client.Jar.Cookies(u))
	// cookies, err := bihttp.LoadCookies(".temp/cookie.json")
	// check(err)
	// log.Println(cookies)
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
