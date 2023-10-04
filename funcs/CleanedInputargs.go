package funcs

func CleanedInputargs(inputargs []string) []string {
	cleanedInputargs := make([]string, 0, len(inputargs))
	for _, arg := range inputargs {
		if arg != "" {
			cleanedInputargs = append(cleanedInputargs, arg)
		}
	}
	return cleanedInputargs
}
