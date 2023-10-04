package main

import (
	"AMJ/funcs"
	"fmt"
	"os"
	"strings"
)

func main() {
	//1. Checking if we have the right amount of arguments
	if len(os.Args) < 2 || len(os.Args) > 8 {
		fmt.Println(funcs.ErrorsM("Usage"))
		fmt.Println(funcs.ErrorsM("AvColor"))
		return
	}
	inputargs := os.Args[1:]
	filename := "standard.txt"
	coloredLetters := ""
	color := ""
	output := ""
	align := ""
	// reverse := ""
	restoprint := "" // for the output file

	if funcs.AllEmpty(inputargs) {
		os.Exit(0)
	}
	// Checking input arguments have any of those flags (--reverse=<fileName> || --align=<type> || --output=<fileName.txt> )
	for i := 0; i < len(inputargs); i++ {
		if strings.HasPrefix(inputargs[i], "--reverse") {
			// reverse := strings.TrimPrefix(inputargs[i], "--reverse=")
			funcs.ReverseAsciiArt()
			os.Exit(0)
		}
		if strings.HasPrefix(inputargs[i], "--align") {
			align = strings.TrimPrefix(inputargs[i], "--align=")
			inputargs[i] = ""
		}
		if strings.HasPrefix(inputargs[i], "--output") {
			output = strings.TrimPrefix(inputargs[i], "--output=")
			inputargs[i] = ""
		}
	}
	inputargs = funcs.CleanedInputargs(inputargs)
	lastindex := len(inputargs)
	// Check for --color flag , extract colored letters and banner filename
	if len(inputargs) >= 1 {
		str := inputargs[0]

		if strings.HasPrefix(inputargs[0], "--color") {
			color = strings.TrimPrefix(inputargs[0], "--color=")
			if funcs.Color(color) == "" {
				// Color doesn't exists in the map.
				fmt.Println(funcs.ErrorsM("AvColor")) //
				os.Exit(0)
			}
			if lastindex == 4 {
				coloredLetters = inputargs[1]
				str = inputargs[2]
				filename = funcs.Banners(inputargs[3])
			} else if lastindex == 3 {
				if funcs.Banners(inputargs[2]) != "" {
					coloredLetters = "   "
					str = inputargs[1]
					filename = funcs.Banners(inputargs[2])
				} else {
					coloredLetters = inputargs[1]
					str = inputargs[2]
				}
			} else if lastindex == 2 {
				coloredLetters = "   "
				str = inputargs[1]
			} else {
				fmt.Println(funcs.ErrorsM("Usage"))
				fmt.Println(funcs.ErrorsM("AvColor"))
				os.Exit(0)
			}
		} else if lastindex == 2 {
			str = inputargs[0]
			filename = funcs.Banners(inputargs[1])
		} else if len(inputargs) != 1 {
			fmt.Println(funcs.ErrorsM("Usage"))
			fmt.Println(funcs.ErrorsM("AvColor"))
			os.Exit(0)
		}

		str = funcs.Alphabetformat(str) // fixing the \t
		res := ""
		restemp := ""
		args := strings.Split(str, "\\n")
		if funcs.AllEmpty(args) {
			args = args[1:]
		}
		// Writing text line by line into res
		if align == "justify" {
			restoprint = funcs.Justify(args, filename, color, coloredLetters, align, output)
		} else {
			for _, word := range args {
				if word == "" {
					fmt.Println()
					restoprint = restoprint + "\n"
					continue
				}

				for i := 0; i < 8; i++ {
					res, restemp = funcs.WhatToPrint(i, word, filename, color, coloredLetters, align, res, restemp)
					if align != "" {
						space := funcs.AlignString(restemp, align, args)
						res = space + res
						fmt.Println(res)
					} else {
						fmt.Println(res)
					}
					restoprint = restoprint + "\n" + restemp
					res = ""
					restemp = ""
				}
			}
		}
	} else {
		fmt.Println(funcs.ErrorsM("Usage"))
		fmt.Println(funcs.ErrorsM("AvColor"))
		return
	}
	if output != "" {
		err := funcs.WriteToFile(restoprint[1:], output)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		}
	}
	os.Exit(0)
}
