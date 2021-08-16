package restful

import (
	"fmt"

	"github.com/bestpilotingalaxy/fbs-test-case/internal/math"
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
	fmt.Println(start)
	res := math.FibonacciBig(start, end)
	log.Infof("start: %d, end: %d", start, end)
	return c.JSON(res)
}
