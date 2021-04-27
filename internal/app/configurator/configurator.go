package configurator

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"webserver/internal/app/storage"
)

type Config struct {
	DBConfig storage.Config `json:"storage"`
}

func Open(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	fContent, err := ioutil.ReadAll(f)

	if err != nil {
		return nil, err
	}

	var cfg Config
	json.Unmarshal(fContent, &cfg)

	defer f.Close()
	return &cfg, nil
}
