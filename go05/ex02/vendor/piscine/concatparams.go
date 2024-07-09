package piscine

func ConcatParams(args []string) string {
	result := make([]string, 0)
	for i, arg := range args {
		if i > 0 {
			result = append(result, "\n")
		}
		result = append(result, arg)
	}
	return joinString(result)
}

func joinString(s []string) string {
	var output string
	for _, str := range s {
		output += str
	}
	return output
}
