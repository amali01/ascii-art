package funcs

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ConsoleWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(ErrorsM("size"))
		os.Exit(1)
	}

	outStr := string(out)
	outStr = strings.TrimSpace(outStr)
	heightWidth := strings.Split(outStr, " ")
	width, err := strconv.Atoi(heightWidth[1])
	if err != nil {
		fmt.Println(ErrorsM("size"))
		os.Exit(1)
	}

	return (width)
}
