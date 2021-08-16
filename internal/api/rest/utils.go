package rest

import (
	"fmt"
	"strconv"
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
