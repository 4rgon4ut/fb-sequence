package restful

import (
	"fmt"

	"github.com/bestpilotingalaxy/fbs-test-case/internal/math"
	"github.com/bestpilotingalaxy/fbs-test-case/internal/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	log "github.com/sirupsen/logrus"
)

// FibonacciSliceHandler ...
func FibonacciSliceHandler(c *fiber.Ctx) error {
	start, err := formatQueryParam(utils.CopyString(c.Query("start")))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	end, err := formatQueryParam(utils.CopyString(c.Query("end")))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if start > end {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid range")
	}

	cachedSequence, err := redis.Client.ZRangeByScore(fmt.Sprint(start), fmt.Sprint(end))
	if err != nil {
		log.Error(err)
	}
	if len(cachedSequence) == int(end-start+1) {
		log.Debug("Cache hit")
		res := formatRedisRecordsToPayload(cachedSequence)
		return c.JSON(res)

	}
	res := math.FibonacciBig(start, end)
	go redis.Client.ZAddNXMany(res)

	return c.JSON(res)
}
