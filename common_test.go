package algorithm

import "testing"

func Test_Reverse_Integer(t *testing.T) {
	result := reverse(123)
	expected := 321
	if result != expected {
		t.Fatalf("Expected: %d | Actual: %d", expected, result)
	}
}
