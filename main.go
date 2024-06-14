package main

import (
	"biliget/biliinfo"
	"biliget/biliinfo/bvinfo"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	bytes, err := biliinfo.BiliGetAddress("BV1YJ4m1g7EQ")
	check(err)
	err = os.WriteFile(".temp/debug.html", bytes, 0666)
	check(err)
	info, err := bvinfo.ParseBvInfo(string(bytes))
	check(err)
	if len(info.Videos) > 0 {
		bytes, err := biliinfo.BiliGetUrl(info.Videos[0].UrlBackup[0])
		check(err)
		err = os.WriteFile(".temp/video.mp4", bytes, 0666)
		check(err)
		log.Printf("Download video (%d, %d)", info.Videos[0].Width, info.Videos[0].Height)
	}
	if len(info.Audios) > 0 {
		bytes, err := biliinfo.BiliGetUrl(info.Audios[0].UrlBackup[0])
		check(err)
		err = os.WriteFile(".temp/audio.mp4", bytes, 0666)
		check(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
