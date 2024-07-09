package piscine

func AppendRange(min, max int) []int {
	var result []int // nil
	for n := min; n < max; n++ {
		result = append(result, n)
	}
	return result
}
