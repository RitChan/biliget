// 1. All config are stored in a json file.
//
// 2. When initConfig(), config will be loaded into BiligetConfig and can be accessed via Global().
//
// 3. User can specify config fileapth through BILIGET_CONFIG env variable or "-c" "--config" cli args.
// If no path provide or the file does not exist, default values will be set.
//
// 4. There is a Filepath field in BiligetConfig which will be the path provided in 3 (if any).
// Filepath will be set even if the provided file don't exist.
package config

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type BiligetConfig struct {
	Filepath    string         // Manual reflection :)
	CookieCache []*http.Cookie `json:"cookie_cache"`
}

var global *BiligetConfig

func Global() *BiligetConfig {
	if global == nil {
		initConfig()
	}
	return global
}

func Dump(filepath string) error {
	if global == nil {
		log.Println("Warning: global config is nil, no file saved")
		return nil
	}
	bytes, err := json.Marshal(Global())
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, bytes, 0666)
	if err != nil {
		return err
	}
	return nil
}

func initConfig() *BiligetConfig {
	global = new(BiligetConfig)
	err := loadFile(global, getConfigPath())
	if err != nil {
		log.Printf("Warning: %s\n", err.Error())
		setDefault(global)
	}
	return global
}

func setDefault(config *BiligetConfig) {
	config.CookieCache = make([]*http.Cookie, 0)
	config.Filepath = ""
}

func loadFile(config *BiligetConfig, filepath string) error {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, config)
	if err != nil {
		return err
	}
	config.Filepath = filepath
	return nil
}

func getConfigPath() string {
	// CLI
	argv := os.Args[1:]
	configFlag := false
	for _, arg := range argv {
		if configFlag {
			return arg
		}
		if arg == "-c" || arg == "--config" {
			configFlag = true
		}
	}
	// Env
	configPath := os.Getenv("BILIGET_CONFIG")
	if configPath != "" {
		return configPath
	}
	// Default
	return "biliget.json"
}
