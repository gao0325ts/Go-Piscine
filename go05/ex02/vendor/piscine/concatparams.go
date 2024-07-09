package piscine

func ConcatParams(args []string) string {
	parts := make([]string, 0)
	for i, arg := range args {
		if i > 0 {
			parts = append(parts, "\n")
		}
		parts = append(parts, arg)
	}
	return joinStrings(parts)
}

func joinStrings(s []string) string {
	var result string
	for _, str := range s {
		result += str
	}
	return result
}
