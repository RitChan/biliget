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
	"path"
)

type BiligetConfig struct {
	DownloadDirectory string         `json:"download_directory"`
	CookieCache       []*http.Cookie `json:"cookie_cache"`
}

var global *BiligetConfig

func Global() *BiligetConfig {
	if global == nil {
		initConfig()
	}
	return global
}

func Dump(config *BiligetConfig, filepath string) error {
	if config == nil {
		log.Println("Warning: config is nil, no file saved")
		return nil
	}
	bytes, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, bytes, 0666)
	if err != nil {
		return err
	}
	return nil
}

func DefaultPath() (string, error) {
	p := userConfigPath()
	if p != "" {
		return p, nil
	}
	p, err := appConfigPath()
	if err != nil {
		return "", err
	}
	return p, nil
}

// Tolerate any error
func initConfig() *BiligetConfig {
	global = new(BiligetConfig)
	setDefault(global)
	p, err := DefaultPath()
	if err != nil {
		return global
	}
	err = loadFile(global, p)
	if err != nil {
		setDefault(global)
	}
	return global
}

// Ensure all fields are set
func setDefault(config *BiligetConfig) {
	config.CookieCache = make([]*http.Cookie, 0)
	config.DownloadDirectory = ""
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
	return nil
}

func userConfigPath() string {
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
	return ""
}

func appConfigPath() (string, error) {
	p, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return path.Join(p, "biliget", "biliget.json"), nil
}
