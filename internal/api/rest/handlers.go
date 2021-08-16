package rest

import (
	"github.com/bestpilotingalaxy/fbs-test-case/internal/math"
	"github.com/bestpilotingalaxy/fbs-test-case/internal/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
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

	res, err := redis.GetFromCache(start, end)
	if err != nil {
		calculatedRes, toCache := math.FibonacciBig(start, end)
		go redis.Client.ZAddNXMany(toCache)
		return c.JSON(calculatedRes)
	}
	return c.JSON(res)
}
