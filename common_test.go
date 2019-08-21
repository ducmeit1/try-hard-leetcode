package algorithm

import "testing"

func Test_Two_Sum(t *testing.T) {
	nums := []int{2,7,14,19}
	target := 9
	expected := []int{0, 1}
	result := twoSum(nums, target)
	for i := 0; i < len(expected); i++ {
		if expected[i] != result[i] {
			t.Fatalf("Expected: %v  | Actual: %v", expected, result)
		}
	}
}

func Test_Reverse_Integer(t *testing.T) {
	result := reverse(-987)
	expected := -789
	if result != expected {
		t.Fatalf("Expected: %d | Actual: %d", expected, result)
	}
}
