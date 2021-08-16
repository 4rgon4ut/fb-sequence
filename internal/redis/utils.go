package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

// GetFromCache ...
func GetFromCache(start uint64, end uint64) (map[uint64]string, error) {
	cachedSequence, err := Client.ZRangeByScore(fmt.Sprint(start), fmt.Sprint(end))
	if err != nil {
		log.Error(err)
	}
	if len(cachedSequence) == int(end-start+1) {
		log.Debug("Cache hit")
		res := formatRedisRecordsToPayload(cachedSequence)
		return res, nil

	}
	return nil, fmt.Errorf("Sequence not in cache")
}

func formatRedisRecordsToPayload(records []redis.Z) map[uint64]string {
	result := make(map[uint64]string)
	for _, v := range records {
		result[uint64(v.Score)] = v.Member.(string)
	}
	return result
}
