package main

import (
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	"github.com/bestpilotingalaxy/fbs-test-case/config"
	"github.com/bestpilotingalaxy/fbs-test-case/internal/api/protobuff"
	"github.com/bestpilotingalaxy/fbs-test-case/internal/api/restful"
	"github.com/bestpilotingalaxy/fbs-test-case/internal/redis"
)

func main() {
	cfg := config.New()
	config.SetupLogger(cfg.LogLevel)

	redis.Client = redis.New(&cfg.Redis)
	api := restful.NewRouter(&cfg.RESTServer)
	go api.RunAPI()

	grpcSrv := protobuff.NewServer(&cfg.GRPCServer)
	protobuff.RunServer(grpcSrv, cfg.GRPCServer.Port)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.wstools.ChatHub.BroadcastToAll
	<-c

	api.Shutdown()
	grpcSrv.Stop()

	log.Info("Good bye!")
	os.Exit(0)
}
