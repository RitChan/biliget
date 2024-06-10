package main

import (
	"biliget/src/biliinfo"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	raw_string, err := biliinfo.GetDocument("BV1nz421h76u")
	check(err)
	info, err := biliinfo.ParseBvInfo(raw_string)
	check(err)
	log.Println(info)
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
