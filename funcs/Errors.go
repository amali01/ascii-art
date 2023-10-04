package funcs

func ErrorsM(err string) string {
	var errMessages = map[string]string{
		"Usage":      "\nUsage: go run main.go followed by one of these options as separate arguments\n\n*in the same order*\n1- [STRING]\n2- [STRING] [BANNER]\n3- --color=<color> [STRING]\n4- --color=<color> [STRING] [BANNER]\n5- --color=<color> <letters to be colored> [STRING]\n6- --color=<color> <letters to be colored> [STRING] [BANNER]\n\n*in any order*\n7- --output=<fileName.txt>\n8- --align=<type>  ( left right center justify )\n\n**********\nto work it has to be actvited separately\n 9- --reverse=<fileName>",
		"AvColor":    "\nThose are the available colors:\nred, green, yellow, blue, purple, cyan, white, gray, darkred, orange, pink,\ngold, teal, lavender, brown, lightblue, magenta, olive, salmon, skyblue, darkpurple, lime\n",
		"banner":     "The Banner input is not available\n choose one of those: shadow, standard, thinkertoy, doom.",
		"align":      "Wrong Alignment input.\n use one of those options:\n( center, left, right, justify ).",
		"fCont":      "Can't read the Banner file (standard.txt), make sure the filename is correct.",
		"revUsage":   "Usage: go run . [OPTION]\n\nEX: go run main.go --reverse=<fileName>",
		"textCont":   "Can't read the input file , make sure the filename is correct",
		"fit":        "The text dose not fit the terminal size.",
		"size":       "Error measuring console size:",
		"Usage1":     "\nUsage: go run main.go [STRING]",
		"Usagefs":    "\nUsage: go run main.go followed by one of these options\n\n*in the same order*\n1- [STRING]\n2- [STRING] [BANNER]",
		"UsageColor": "\nUsage: go run main.go followed by one of these options\n\n*in the same order*\n1- [STRING]\n2- [STRING] [BANNER]\n3- --color=<color> [STRING]\n4- --color=<color> [STRING] [BANNER]\n5- --color=<color> <letters to be colored> [STRING]\n6- --color=<color> <letters to be colored> [STRING] [BANNER]",
		"UsageAlign": "\nUsage: go run main.go followed by one of these options\n\n*in the same order*\n1- [STRING]\n2- [STRING] [BANNER]\n3- --color=<color> [STRING]\n4- --color=<color> [STRING] [BANNER]\n5- --color=<color> <letters to be colored> [STRING]\n6- --color=<color> <letters to be colored> [STRING] [BANNER]\n\n*in any order*\n7- --align=<type> ",
		"UsageOut":   "\nUsage: go run main.go followed by one of these options\n\n*in the same order*\n1- [STRING]\n2- [STRING] [BANNER]\n3- --color=<color> [STRING]\n4- --color=<color> [STRING] [BANNER]\n5- --color=<color> <letters to be colored> [STRING]\n6- --color=<color> <letters to be colored> [STRING] [BANNER]\n\n*in any order*\n7- --output=<fileName.txt>\n8- --align=<type>  ( left right center justify )",
		// "UsageOut": "asdfghf",
		// Add more error Messages here as needed
	}

	if error1, ok := errMessages[err]; ok {
		return error1
	}
	return ("")
}
