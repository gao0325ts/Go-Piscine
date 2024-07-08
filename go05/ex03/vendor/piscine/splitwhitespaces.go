package piscine

func SplitWhiteSpaces(s string) []string {
	result := make([]string, 0)
	var tmp string
	var i int
	for _, r := range s {
		if r == ' ' {
			result = append(result, tmp)
			i++
			tmp = ""
			continue
		}
		tmp += string(r)
	}
	result = append(result, tmp)
	return result
}
