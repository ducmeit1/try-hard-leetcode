package algorithm

import "math"

func isPalindrome(x int) bool {
	arr := make([]int, 0)
	sign := getSign(x)
	tmp := x
	x = int(math.Abs(float64(x)))
	for x > 0 {
		ld := x % 10
		if x * sign == tmp {
			ld = ld * sign
		}
		arr = append(arr, ld)
		x = x / 10
	}

	length := len(arr)
	if length == 1 {
		if arr[0] >= 0 {
			return true
		}
		return false
	}

	if length == 2 {
		if arr[0] != arr[1] {
			return false
		}
		return true
	}

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

func getSign(x int) int {
	if x > 0 {return 1}
	return -1
}
