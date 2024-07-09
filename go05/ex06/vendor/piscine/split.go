package piscine

func Split(s, sep string) []string {
	sLen := strLen(s)
	sepLen := strLen(sep)
	var start int
	result := make([]string, 0)

	for i := 0; i <= sLen-sepLen; i++ {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			start = i + sepLen
		}
	}

	if start <= sLen {
		result = append(result, s[start:])
	}

	return result
}

func strLen(s string) int {
	var length int
	for range s {
		length++
	}
	return length
}