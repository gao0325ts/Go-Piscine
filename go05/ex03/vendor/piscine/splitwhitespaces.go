package piscine

func SplitWhiteSpaces(s string) []string {
	result := make([]string, 0)
	var start int
	var end int
	for i, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if start == i {
				result = append(result, "")
			} else {
				result = append(result, s[start:end])
			}
			start = i + 1
		}
		end++
	}
	if start == end {
		result = append(result, "")
	} else if start < end {
		result = append(result, s[start:end])
	}
	return result
}