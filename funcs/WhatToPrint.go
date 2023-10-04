package funcs

import (
	"fmt"
	"os"
	"strings"
)

func WhatToPrint(i int, word string, filename string, color string, coloredLetters string, align string, res string, restemp string) (string, string) {
	count := 0
	for k, letter := range word {
		line := GetLine(1+(int(letter)-32)*9+i, filename)
		restemp += line
		// Apply color to the specified letters
		if strings.ContainsRune(coloredLetters, letter) {
			if k < (1 + len(word) - len(coloredLetters)) {
				if coloredLetters == word[k:k+len(coloredLetters)] {
					count = len(coloredLetters)
				}
			}
			if count != 0 {
				// Check if the color exists in the colors map
				escapeSeq := Color(color)
				if escapeSeq != "" {
					// Color exists in the map, do something
					line = escapeSeq + line + Color("reset")
				} else {
					fmt.Println(ErrorsM("AvColor"))
					os.Exit(0)
				}
				count--
			}
		} else if coloredLetters == "   " {
			// Check if the color exists in the colors map
			escapeSeq := Color(color)
			if escapeSeq != "" {
				// Color exists in the map, do something
				line = escapeSeq + line + Color("reset")
			} else {
				fmt.Println(ErrorsM("AvColor"))
				os.Exit(0)
			}
		}
		res += line
		// Optionally, you can adjust the spacing between characters here
	}
	return res, restemp
}
