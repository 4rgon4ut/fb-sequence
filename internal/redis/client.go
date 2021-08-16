package redis

import (
	"context"
	"fmt"

	"github.com/bestpilotingalaxy/fbs-test-case/config"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

// Client ...
var Client *Redis

// Redis ...
type Redis struct {
	Client  *redis.Client
	Config  *config.Redis
	Context context.Context
}

// New ...
func New(c *config.Redis) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     ":" + c.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Redis{
		Client:  client,
		Config:  c,
		Context: context.Background(),
	}
}

// GET ...
func (r *Redis) GET(key string) (string, error) {
	ctx, cancel := context.WithTimeout(r.Context, 10)
	defer cancel()

	val, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("no such key")
	} else if err != nil {
		return "", fmt.Errorf("cant GET key '%s' from redis: %v", key, err)
	}
	return val, nil
}

// SET ...
func (r *Redis) SET(key string, val string) error {
	ctx, cancel := context.WithTimeout(r.Context, 10)
	defer cancel()
	err := r.Client.Set(ctx, key, val, 0).Err()
	if err != nil {
		return fmt.Errorf("cant SET <%s : %s> to redis: %v", key, val, err)
	}
	return nil
}

// ZADDNXMany ...
func (r *Redis) ZADDNXMany(result map[uint64]string) error {
	addSlice := make([]*redis.Z, 0)
	for k, v := range result {
		member := &redis.Z{Score: float64(k), Member: v}
		addSlice = append(addSlice, member)
	}
	ctx, cancel := context.WithTimeout(r.Context, 10)
	defer cancel()
	err := r.Client.ZAddNX(ctx, r.Config.SetName, addSlice...).Err()
	if err != nil {
		err := fmt.Errorf("cant ZADDNX sequence to redis: %v", err)
		log.Error(err.Error())
		return err
	}
	return nil
}
