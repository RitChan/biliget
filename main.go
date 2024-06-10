package main

import (
	"biliget/src/biliinfo"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	raw, err := os.ReadFile(".temp/rnbvocal.html")
	raw_string := string(raw)
	check(err)
	info, err := biliinfo.ParseRawInitialState(raw_string)
	check(err)
	log.Println(info)
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
