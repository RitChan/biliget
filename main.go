package main

import (
	"biliget/src/biliinfo"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	bytes, err := biliinfo.BiliGetAddress("BV1YJ4m1g7EQ")
	check(err)
	info, err := biliinfo.ParseBvInfo(string(bytes))
	check(err)
	log.Println(info)
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
