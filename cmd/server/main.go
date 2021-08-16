package main

import (
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	"github.com/bestpilotingalaxy/fbs-test-case/config"
	"github.com/bestpilotingalaxy/fbs-test-case/internal/api/protobuff"
)

func main() {
	cfg := config.New()
	config.SetupLogger(cfg.LogLevel)

	// api := restful.NewRouter(&cfg.Server)
	// api.RunAPI()

	grpcSrv := protobuff.NewServer()
	protobuff.RunServer(grpcSrv)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.wstools.ChatHub.BroadcastToAll
	<-c

	log.Info("Good bye!")
	os.Exit(0)
}
