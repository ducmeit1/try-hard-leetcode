package algorithm

import (
	"math"
)

func reverseNumber(x int) int {
	result := 0
	for x > 0 {
		ld := x % 10
		result = result * 10 + ld
		x = x/10
	}

	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}

	return result
}

func reverse(x int) int {
	signed := 1
	if x < 0 {
		signed = -1
	}
	x = int(math.Abs(float64(x)))
	return reverseNumber(x) * signed
}