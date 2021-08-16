package math

import (
	"math/big"
)

// FibonacciBig ...
func FibonacciBig(start uint64, end uint64) (map[uint64]string, map[uint64]string) {
	result := make(map[uint64]string)
	sequenceToCache := make(map[uint64]string)
	if end <= 1 {
		result[end] = big.NewInt(int64(end)).String()
	}
	var n1, n2 = big.NewInt(0), big.NewInt(1)

	for i := uint64(1); i <= end; i++ {
		n1.Add(n1, n2)
		n1, n2 = n2, n1
		if i >= start {
			result[i] = n1.String()
		}
		sequenceToCache[i] = n1.String()
	}

	return result, sequenceToCache
}
