package redis

// DigitLength return string length by digits (not by bytes)
func DigitLength(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}
