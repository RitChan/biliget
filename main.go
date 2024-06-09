package main

import (
	"biliget/src/biliinfo"
	"log"
	"os"
)

func main() {
	// log.SetOutput(os.Stdout)
	raw, err := os.ReadFile(".temp/rnbvocal.html")
	check(err)
	info, err := biliinfo.ParseRawPlayInfo(string(raw))
	check(err)
	log.Println(info)
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
