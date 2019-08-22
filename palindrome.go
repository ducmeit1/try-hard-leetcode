package algorithm

import "math"

func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x > 0) {
		return false
	}

	arr := make([]int, 0)
	x = int(math.Abs(float64(x)))
	for x > 0 {
		ld := x % 10
		arr = append(arr, ld)
		x = x / 10
	}

	length := len(arr)

	mid := length / 2
	s1 := arr[0:mid]
	s2 := arr[mid:length]

	n := len(s1)
	l := len(s2)
	for i:=0;i< n;i++ {
		if s1[i] != s2[l - 1] {
			return false
		}
		l--
	}

	return true
}