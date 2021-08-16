package restful

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func formatQueryParam(param string) (uint64, error) {
	if param == "" {
		return 0, fmt.Errorf("start/end param not specified")
	}
	num, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("start/end param must be a number")
	}
	return num, nil
}

func formatRedisRecordsToPayload(records []redis.Z) map[int]string {
	result := make(map[int]string)
	for _, v := range records {
		result[int(v.Score)] = v.Member.(string)
	}
	return result
}
