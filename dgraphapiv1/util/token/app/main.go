package main

import (
	"log"

	"github.com/folivoralabs/api/pkg/auth/tokens"
	"github.com/folivoralabs/api/pkg/config"
)

func main() {
	log.Println("start token fetching")

	cfg, err := config.New("../../../etc/config/config.json")
	if err != nil {
		log.Fatal("error reading config file:", err.Error())
	}

	tokensClient := tokens.New(cfg)

	token, err := tokensClient.GetAppToken()
	if err != nil {
		log.Fatal("error fetching token:", err.Error())
	}

	log.Println("token:", token)
}
