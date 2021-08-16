package math

import (
	"fmt"
	"math"
	"math/big"
)

// FibonacciBig ...
func FibonacciBig(start uint64, end uint64) map[uint64]string {
	result := make(map[uint64]string)
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
	}
	return result
}

// Fibonacci ...
func Fibonacci(n uint) map[uint]uint64 {
	result := make(map[uint]uint64)
	if n <= 1 {
		result[n] = uint64(n)
		return result
	}

	var n1, n2 uint64 = 0, 1

	for i := uint(2); i <= n; i++ {
		n1, n2 = n2, n1+n2
		result[i] = n2
	}

	return result
}

// // FibGenerator ...
// func FibGenerator(start *big.Int, end *big.Int) chan int {
// 	c := make(chan *big.Int)

// 	go func() {
// 		for  ; i, j = i+j, i {
// 			c <- i
// 		}
// 	}()

// 	return c
// }
func FibBinet(start int, end int) map[int]string {
	result := make(map[int]string)
	for start := start; start <= end; start++ {
		g := (1 + math.Sqrt(5)) / 2
		ret := (math.Pow(g, float64(start)) - math.Pow(1-g, float64(start))) / math.Sqrt(5)
		r := big.NewFloat(ret)
		num, _ := r.Uint64()
		result[start] = fmt.Sprint(num)
	}
	return result
}

// 	return
// }
