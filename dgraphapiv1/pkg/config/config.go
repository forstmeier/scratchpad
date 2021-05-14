package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config is a struct representation of the root/secret config.json file.
type Config struct {
	Auth0    auth0    `json:"AUTH0"`
	Folivora folivora `json:"FOLIVORA"`
}

type auth0 struct {
	DomainURL   string          `json:"DOMAIN_URL"`
	AudienceURL string          `json:"AUDIENCE_URL"`
	TokenURL    string          `json:"TOKEN_URL"`
	UI          application     `json:"UI"`
	API         application     `json:"API"`
	Users       map[string]user `json:"USERS"`
}

type application struct {
	ClientID     string `json:"CLIENT_ID"`
	ClientSecret string `json:"CLIENT_SECRET"`
}

type user struct {
	ID       string `json:"ID"`
	Username string `json:"USERNAME"`
	Password string `json:"PASSWORD"`
}

type folivora struct {
	CustomURL    string `json:"CUSTOM_URL"`
	CustomSecret string `json:"CUSTOM_SECRET"`
	DgraphURL    string `json:"DGRAPH_URL"`
}

// New reads in the config.json file present in the path variable and generates
// the struct representation.
func New(path string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, newErrorReadFile(err)
	}

	config := &Config{}
	if err := json.Unmarshal(configBytes, config); err != nil {
		return nil, newErrorUnmarshalConfig(err)
	}

	return config, nil
}
