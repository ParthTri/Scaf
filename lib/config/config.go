package config

import (
	"errors"
	"fmt"
	"os"
)

type ConfigDir struct {
	Path			string
	Templates []os.DirEntry
}

func FindConfigDir() (string, error) {
	config_dir, _ := os.UserConfigDir()
	config_dir = fmt.Sprintf("%v/scaf/", config_dir)	
	_, err := os.ReadDir(config_dir)

	if (!os.IsExist(err)) {
		return config_dir, nil
	}

	home_scaf, _ := os.UserHomeDir()
	home_scaf = fmt.Sprint(home_scaf, "/.scaf/")

	_, err = os.ReadDir(home_scaf)

	if (!os.IsExist(err)) {
		return home_scaf, nil
	}

	return "", errors.New("Config files not found")
}