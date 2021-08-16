package config

import (
	"context"
	"os"

	"sync"

	"github.com/sethvargo/go-envconfig"
	log "github.com/sirupsen/logrus"
)

var once sync.Once

// RESTServer configuration
type RESTServer struct {
	Port string `env:"SERVER_PORT"`
}

// GRPCServer ...
type GRPCServer struct {
	Port string `env:"GRPC_SERVER_PORT"`
}

// Redis ...
type Redis struct {
	Port    string `env:"REDIS_PORT"`
	SetName string `env:"REDIS_SETNAME"`
}

// Config global config struct
type Config struct {
	RESTServer
	GRPCServer
	LogLevel string `env:"LOG_LEVEL"`
}

// New creates config struct and fills with env variables
func New() *Config {
	ctx := context.Background()
	c := &Config{}
	process := func() {
		if err := envconfig.Process(ctx, c); err != nil {
			log.Fatal(err)
		}
	}
	once.Do(process)
	return c
}

// SetupLogger ...
func SetupLogger(level string) {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	switch level {
	case "INFO":
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
}
