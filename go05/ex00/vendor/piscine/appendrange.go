package piscine

func AppendRange(min, max int) []int {
	var result []int
	for i := 0; min < max; i++ {
		result = append(result, min)
		min++
	}
	return result
}
