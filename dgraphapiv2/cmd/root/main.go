package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/folivoralabs/api/cmd/root/handler"
	"github.com/folivoralabs/api/cmd/root/server"
	"github.com/folivoralabs/api/pkg/config"
)

func main() {
	rootConfig, err := config.New("../../etc/config/config.json")
	if err != nil {
		log.Fatalf("error reading config: %s\n", err.Error())
	}

	rootClient := handler.NewClient(rootConfig.App.DBURL)

	rootServer := server.New(rootClient.Handler())

	ctx := context.Background()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go rootServer.Start()
	go func() {
		<-sigs
		done <- true
	}()

	<-done
	rootServer.Stop(ctx)
}
