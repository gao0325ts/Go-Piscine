package piscine

func ConcatParams(args []string) string {
	var result string
	for i, p := range args {
		if i != 0 {
			result += "\n"
		}
		result += p
	}
	return result
}
