package config

import (
	"biliget/biliinfo/bihttp"
)

func Initialize() {
	initLogging()
	initConfig()
	initializeCookies()
}

func initializeCookies() {
	bihttp.ClientAddCookies(bihttp.GetClient(), Global().CookieCache)
}
