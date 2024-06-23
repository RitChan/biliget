package config

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
)

type FileConfig struct {
	ConfigWritePath     string
	DownloadDirectories []string
	CookieCache         []*http.Cookie
}

var global *FileConfig

func Global() *FileConfig {
	if global == nil {
		initConfig()
	}
	return global
}

func Dump(config *FileConfig) error {
	if config == nil {
		log.Println("Warning: config is nil, no file saved")
		return nil
	}
	bytes, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile(config.ConfigWritePath, bytes, 0666)
	if err != nil {
		return err
	}
	return nil
}

// Tolerate any error
func initConfig() *FileConfig {
	configA := newDefault()
	p, err := appConfigPath()
	if err != nil {
		global = configA
		return global
	}
	configB, err := newFile(p)
	if err != nil {
		global = configA
		return global
	}
	global = configB
	return global
}

func newDefault() *FileConfig {
	c := new(FileConfig)
	setDefault(c)
	return c
}

func newFile(filepath string) (*FileConfig, error) {
	c := new(FileConfig)
	setDefault(c)
	err := setFile(c, filepath)
	return c, err
}

// Ensure all fields are set
func setDefault(config *FileConfig) {
	config.CookieCache = make([]*http.Cookie, 0)
	config.ConfigWritePath = ""
	config.DownloadDirectories = nil
}

func setFile(config *FileConfig, filepath string) error {
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

func appConfigPath() (string, error) {
	p, err := os.Executable()
	if err != nil {
		return "", err
	}
	d := path.Dir(p)
	return path.Join(d, "biliget.json"), nil
}
