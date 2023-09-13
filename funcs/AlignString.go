package funcs

import (
	"fmt"
	"os"
	"strings"
)

func AlignString(lines string, align string, args []string) string {
	spaceNum := len(args) - 1
	var TerminalWidth int
	TerminalWidth = ConsoleWidth()
	if len(lines) > TerminalWidth {
		fmt.Println(ErrorsM("fit"))
		os.Exit(0)
	}
	if align != "left" && align != "right" && align != "center" && align != "justify" {
		fmt.Println(ErrorsM("align"))
		os.Exit(0)
	}
	var result strings.Builder
	switch align {
	case "left":
		lines = ""
	case "right":
		lines = strings.Repeat(" ", TerminalWidth-len(lines)-1)
	case "center":
		lines = strings.Repeat(" ", (TerminalWidth-len(lines))/2)
	case "justify":
		lines = strings.Repeat(" ", (TerminalWidth-len(lines))/spaceNum)
	}
	result.WriteString(lines)

	return result.String()
}
