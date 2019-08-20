package algorithm

func twoSum(numbers []int, target int) []int {
	m := parseToMap(numbers)
	for i, v := range numbers {
		n1 := v
		n2 := target - n1
		if index := m[n2] - 1; index != -1 && index != i {
			return []int{i, index}
		}
	}
	panic("no result found")
}

func parseToMap(numbers []int) map[int]int {
	m := make(map[int]int)
	for i, v := range numbers {
		m[v] = i + 1
	}
	return m
}
