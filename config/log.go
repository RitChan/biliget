package config

import (
	"log"
	"os"
)

func initLogging() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
