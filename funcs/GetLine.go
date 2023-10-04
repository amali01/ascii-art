package funcs

import (
	"bufio"
	"fmt"
	"os"
)

func GetLine(num int, filename string) string {
	file, e := os.Open(filename)
	if e != nil { // Checking if we have the right file name
		fmt.Println(ErrorsM("banner"))
		os.Exit(0)
	}
	defer file.Close()

	str := ""
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		if lineNumber == num {
			str = scanner.Text()
		}
		lineNumber++
	}
	return str
}
