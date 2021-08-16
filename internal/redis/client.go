package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/bestpilotingalaxy/fbs-test-case/config"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

// Client ...
var Client *Redis

// Redis ...
type Redis struct {
	client  *redis.Client
	Config  *config.Redis
	Context context.Context
}

// New ...
func New(c *config.Redis) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Name + ":" + c.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Redis{
		client:  client,
		Config:  c,
		Context: context.Background(),
	}
}

// ZAddNXMany ...
func (r *Redis) ZAddNXMany(result map[uint64]string) error {
	log.Debug("Saving elements")
	addSlice := make([]*redis.Z, 0)

	for k, v := range result {
		member := &redis.Z{Score: float64(k), Member: v}
		addSlice = append(addSlice, member)
	}

	ctx, cancel := context.WithTimeout(r.Context, 10*time.Second)
	defer cancel()

	err := r.client.ZAddNX(ctx, r.Config.SetName, addSlice...).Err()
	if err != nil {
		err := fmt.Errorf("cant ZADDNX sequence to redis: %v", err)
		log.Error(err.Error())
		return err
	}

	return nil
}

// ZRangeByScore ...
func (r *Redis) ZRangeByScore(start string, end string) ([]redis.Z, error) {
	ctx, cancel := context.WithTimeout(r.Context, 10*time.Second)
	defer cancel()

	rng := &redis.ZRangeBy{Min: start, Max: end}
	vals, err := r.client.ZRangeByScoreWithScores(ctx, r.Config.SetName, rng).Result()

	if err == redis.Nil {
		return nil, fmt.Errorf("no keys in range")
	} else if err != nil {
		return nil, fmt.Errorf("cant get keys range from redis: %v", err)
	}

	return vals, nil
}
