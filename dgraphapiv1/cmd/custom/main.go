package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/folivoralabs/api/cmd/custom/handlers/specimens"
	"github.com/folivoralabs/api/cmd/custom/handlers/users"
	"github.com/folivoralabs/api/cmd/custom/middleware"
	"github.com/folivoralabs/api/cmd/custom/server"
	"github.com/folivoralabs/api/pkg/auth/tokens"
	"github.com/folivoralabs/api/pkg/auth/user"
	"github.com/folivoralabs/api/pkg/config"
)

func main() {
	configPath := flag.String("config", "../../etc/config/config.json", "path to config json file")

	flag.Parse()

	cfg, err := config.New(*configPath)
	if err != nil {
		log.Fatalf("error reading config: %s\n", err.Error())
	}

	tokensClient := tokens.New(cfg)

	appToken, err := tokensClient.GetAppToken()
	if err != nil {
		log.Fatalf("error getting auth0 management api token: %s\n", err.Error())
	}

	root := middleware.New(cfg.Folivora.CustomSecret)

	usersClient := user.New(
		cfg.Auth0.AudienceURL,
		appToken,
	)

	usersHandler := users.Handler(
		cfg.Folivora.DgraphURL,
		tokensClient,
		usersClient,
	)

	specimensHandler := specimens.Handler(
		cfg.Folivora.DgraphURL,
		tokensClient,
	)

	customServer := server.New(root.Middleware, usersHandler, specimensHandler)

	ctx := context.Background()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go customServer.Start()
	go func() {
		<-sigs
		done <- true
	}()

	<-done
	customServer.Stop(ctx)
}
