package funcs

import (
	"fmt"
	"strings"
)

func Justify(args []string, filename string, color string, coloredLetters string, align string, output string) string {
	resToOutput := ""
	tempOutput := ""
	res := ""
	restemp := ""
	space := ""
	var res1 []string
	var tmpLine string

	for _, wordsByLine := range args {
		wordByWord := strings.Split(wordsByLine, " ")
		wordByWord = CleanedInputargs(wordByWord)
		for i := 0; i < 8; i++ {

			for _, word := range wordByWord {
				res, restemp = WhatToPrint(i, word, filename, color, coloredLetters, align, res, restemp)
				if len(wordByWord) == 1 {
					space = AlignString(restemp, "left", wordByWord)
					res = space + res
					fmt.Println(res)
				}
				if i == 0 {
					tmpLine = tmpLine + restemp
				}
				res1 = append(res1, res) // Append the result to the res1 slice
				res = ""
				restemp = ""
			}
			if output != "" {
				_, tempOutput = WhatToPrint(i, wordsByLine, filename, color, coloredLetters, align, res, restemp)
			}
			resToOutput = resToOutput + "\n" + tempOutput
			if len(wordByWord) > 1 {
				space = AlignString(tmpLine, align, wordByWord)
				resWithSpase := strings.Join(res1, space)
				fmt.Println(resWithSpase)
			}
			res1 = nil
		}
		tmpLine = ""
	}

	return resToOutput
}
