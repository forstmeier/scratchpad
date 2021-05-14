package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config is a struct representation of the root/secret config.json file.
type Config struct {
	API api `json:"API"`
}

type api struct {
	ESURL string `json:"ES_URL"`
}

// New reads in the config.json file present in the path variable and generates
// the struct representation.
func New(path string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, ErrorReadFile{err: err}
	}

	config := &Config{}
	if err := json.Unmarshal(configBytes, config); err != nil {
		return nil, ErrorUnmarshalConfig{err: err}
	}

	return config, nil
}
