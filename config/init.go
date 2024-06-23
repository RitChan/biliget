package config

import (
	"biliget/biliinfo/bihttp"
	"biliget/ffmpeg"
)

func Initialize() {
	initLogging()
	initConfig()
	initializeCookies()
	// ignore ffmpeg error, ffmpeg is optional
	// ffmpeg is only used in video/audio processing
	// and this app could have many other functionalities
	ffmpeg.InitFFmpeg()
}

func initializeCookies() {
	bihttp.ClientSetCookies(bihttp.GetClient(), Global().CookieCache)
}
